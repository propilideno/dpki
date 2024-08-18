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
              Delete Domain
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
                label="Domain *"
                :rules="rules.domain"
                debounce="1000"
                :loading="domainLoading"
              />
            </div>
            <div v-if="domainData.exists" class="col-12">
              <q-input v-model="domainData.hash" label="Hash *" />
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
            label="Delete"
            color="negative"
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
import { requiredRule } from "src/utils/rules";

const DEFAULT_DOMAIN_DATA = {
  domain: "",
  hash: "",
  exists: false,
};

const useRules = () => {
  return {
    domain: [requiredRule()],
  };
};

export default defineComponent({
  name: "HashDialog",
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

const domainLoading = ref(false);
const getDomain = async () => {
  domainData.value.exists = false;

  try {
    domainLoading.value = true;

    const { data } = await api.get(`dns/${domainData.value.domain}`);

    domainData.value.exists = true;

    Object.assign(domainData.value, data);
  } finally {
    domainLoading.value = false;
  }
};

const onUpdateDomain = async (newDomain) => {
  domainData.value.domain = newDomain;

  await getDomain();
};

const onSubmit = async () => {
  if (!domainData.value.exists) {
    $q.notify({
      type: "negative",
      message: "This domain doesn't exists.",
    });
    return;
  }
  try {
    loading.value = true;

    await api.patch(`dns/${domainData.value.domain}/clear-txt`, null, {
      headers: {
        "X-Requested-With": "XMLHttpRequest",
        "Content-Type": "Application/json",
        hash: domainData.value.hash,
      },
    });

    $q.notify({
      type: "positive",
      message: "TXT from domain sucessfully deleted.",
    });

    onOKClick();
  } finally {
    loading.value = false;
  }
};
</script>
