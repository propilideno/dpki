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
              Certificate Status
            </div>
          </div>
        </q-card-section>
        <q-separator />
        <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <q-input
                @update:model-value="onUpdateCertificate"
                :model-value="certificate"
                label="Certificate *"
                debounce="1000"
                :rules="rules.certificate"
                autogrow
                filled
              >
                <template #append>
                  <q-icon v-if="statusLoaded" v-bind="getIconProps">
                    <q-tooltip class="text-subtitle2">
                      {{ getIconProps.tooltip }}
                    </q-tooltip>
                  </q-icon>
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
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, ref, reactive, computed } from "vue";
import { useDialogPluginComponent, copyToClipboard, useQuasar } from "quasar";
import { blockchain } from "src/boot/axios";
import { requiredRule } from "src/utils/rules";

const useRules = () => {
  return {
    certificate: [requiredRule()],
  };
};

export default defineComponent({
  name: "CertificateStatusDialog",
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

const status = ref(false);
const certificate = ref("");
const loading = ref(false);
const statusLoaded = ref(false);

const getIconProps = computed(() => {
  return status.value
    ? { name: "check", color: "positive", tooltip: "Certificate is valid" }
    : { name: "close", color: "negative", tooltip: "Certificate is invalid" };
});

const getCertificateStatus = async () => {
  status.value = false;
  statusLoaded.value = false;

  try {
    console.log('teste')
    loading.value = true;

    const { data } = await blockchain.get(`certificate/status/${certificate.value}`);
    status.value = data.status ?? false;
  } finally {
    loading.value = false;
    statusLoaded.value = true;
  }
};

const onUpdateCertificate = async (newCertificate) => {
  certificate.value = newCertificate

  await getCertificateStatus();
};

</script>
