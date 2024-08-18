<template>
  <q-dialog
    ref="dialogRef"
    @hide="onDialogHide"
    @shake="onCloseDialog"
    :position="$q.screen.gt.xs ? 'standard' : 'bottom'"
  >
    <q-card class="medium-dialog">
      <q-form @submit="onSubmit">
        <q-card-section class="row no-wrap">
          <div class="full-width-flex">
            <div class="text-h6 text-grey-9 flex items-center full-height">
              Manage Domain
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
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="domainData.txt"
                label="TXT *"
                :rules="rules.txt"
                :disable="domainLoading"
              />
            </div>
            <div v-if="domainData.exists" class="col-12">
              <q-input
                v-model="domainData.hash"
                label="Hash *"
                :rules="rules.hash"
              />
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
          <q-btn
            type="submit"
            :label="submitLabel"
            color="primary"
            text-color="white"
            padding="sm md"
            :loading="loading"
          />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, ref, reactive } from "vue";
import { useDialogPluginComponent, useQuasar } from "quasar";
import cloneDeep from "lodash.clonedeep";
import { api } from "src/boot/axios";
import ShowSharedKey from "./ShowSharedKey.vue";
import { requiredRule } from "src/utils/rules";

const DEFAULT_DOMAIN_DATA = {
  domain: "",
  txt: "",
  hash: "",
  exists: false,
};

const useRules = () => {
  return {
    domain: [requiredRule()],
    txt: [requiredRule()],
    hash: [requiredRule()],
  };
};

export default defineComponent({
  name: "RegisterDomainDialog",
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
const loading = ref(false);
const submitLabel = ref("Register");

const domainLoading = ref(false);
const getDomain = async () => {
  domainData.value.exists = false;

  try {
    domainLoading.value = true;

    const { data } = await api.get(`dns/${domainData.value.domain}`, {
      customErrorHandlers: {
        404: () => {
          $q.notify({
            type: 'negative',
            message: 'Inexistent domain, register a new one.'
          })
        },
      },
    });

    $q.notify({
      type: 'positive',
      message: 'Edit an existent domain.',
    })

    domainData.value.exists = true;
    submitLabel.value = "Edit";

    Object.assign(domainData.value, data);
  } catch (e) {
    submitLabel.value = "Register";
  } finally {
    domainLoading.value = false;
  }
};

const onUpdateDomain = async (newDomain) => {
  domainData.value.domain = newDomain;

  await getDomain();
};

const onSubmit = async () => {
  try {
    loading.value = true;

    const { data } = await api.post("dns", domainData.value, {
      headers: {
        "X-Requested-With": "XMLHttpRequest",
        "Content-Type": "Application/json",
        hash: domainData.value.hash,
      },
    });

    if (data.shared_key) {
      $q.dialog({
        component: ShowSharedKey,
        componentProps: {
          sharedKey: data.shared_key,
        },
      });
    }

    $q.notify({
      type: "positive",
      message: "Domain sucessfully registered.",
    });

    onOKClick();
  } finally {
    loading.value = false;
  }
};
</script>
