<template>
  <q-dialog
    ref="dialogRef"
    @hide="onDialogHide"
    @shake="onCloseDialog"
    :position="$q.screen.gt.xs ? 'standard' : 'bottom'"
  >
    <q-card class="medium-dialog">
      <q-card-section class="row no-wrap">
        <div class="full-width-flex">
          <div class="text-h6 text-grey-9 flex items-center full-height">
            Show Register Domain
          </div>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <div class="row q-col-gutter-md">
          <div class="col-12">
            <q-input
              @update:model-value="onUpdateDomain"
              :model-value="domainData.domain"
              :rules="rules.domain"
              label="Domain *"
              debounce="1000"
              :loading="domainLoading"
            >
              <template #append>
                <q-icon v-if="domainLoaded" v-bind="getIconProps">
                  <q-tooltip class="text-subtitle2">
                    {{ getIconProps.tooltip }}
                  </q-tooltip>
                </q-icon>
              </template>
            </q-input>
          </div>
          <div v-if="domainData.exists" class="col-12">
            <q-input
              v-model="domainData.txt"
              label="TXT"
              readonly
              autogrow
              filled
              stack-label
            >
              <template #append>
                <q-btn
                  @click="copyTxtToClipboard(domainData.txt)"
                  icon="content_copy"
                  size="0.5em"
                  round
                >
                  <q-tooltip class="text-subtitle2"> Copy TXT </q-tooltip>
                </q-btn>
              </template>
            </q-input>
          </div>
        </div>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn
          @click="onCloseDialog"
          label="Close"
          color="grey-9"
          padding="sm md"
          flat
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, ref, computed, reactive } from "vue";
import { copyToClipboard, useDialogPluginComponent, useQuasar } from "quasar";
import cloneDeep from "lodash.clonedeep";
import { api } from "src/boot/axios";
import { requiredRule } from "src/utils/rules";

const DEFAULT_DOMAIN_DATA = {
  domain: "",
  txt: "",
  hash: "",
  exists: false,
};

const useRules = () => {
  return {
    message: [requiredRule()],
  };
};

export default defineComponent({
  name: "ShowRegisterDomainDialog",
});
</script>

<script setup>
/* const props = */ defineProps({});

/* const emit = */ defineEmits([...useDialogPluginComponent.emits]);

const {
  dialogRef,
  onDialogHide,
  onDialogOK: onOKClick,
  onDialogCancel: onCancelClick,
} = useDialogPluginComponent();

const onCloseDialog = () => {
  onCancelClick();
};

const $q = useQuasar();

const rules = reactive(useRules());

const domainData = ref(cloneDeep(DEFAULT_DOMAIN_DATA));

const domainLoading = ref(false);
const domainLoaded = ref(false);
const getDomain = async () => {
  domainData.value.exists = false;
  domainLoaded.value = false;

  try {
    domainLoading.value = true;

    const { data } = await api.get(`dns/${domainData.value.domain}`, {
      customErrorHandlers: {
        404: () => {
          $q.notify({
            type: 'negative',
            message: 'Inexistent domain.'
          })
        }
      }
    });

    domainData.value.exists = true;

    Object.assign(domainData.value, data);
  } finally {
    domainLoading.value = false;
    domainLoaded.value = true;
  }
};

const onUpdateDomain = async (newDomain) => {
  domainData.value.domain = newDomain;

  await getDomain();
};

const getIconProps = computed(() => {
  return domainData.value.exists
    ? { name: "check", color: "positive", tooltip: "Domain exists" }
    : { name: "close", color: "negative", tooltip: "Domain doesn't exists" };
});

const copyTxtToClipboard = (value) => {
  copyToClipboard(value)
    .then(() => {
      $q.notify({
        type: "positive",
        message: "Copied!",
      });
    })
    .catch(() => {
      $q.notify({
        type: "alert",
        message: "Error on Copy.",
      });
    });
};
</script>
