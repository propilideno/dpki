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
              Request Certificate
            </div>
          </div>
        </q-card-section>
        <q-separator />
        <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <q-input
                v-model="certificateData.domain"
                label="Domain *"
                :rules="rules.domain"
                filled
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="certificateData.certificate"
                label="Certificate *"
                :rules="rules.certificate"
                filled
                autogrow
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
            label="Request"
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
import { useDialogPluginComponent, copyToClipboard, useQuasar } from "quasar";
import cloneDeep from "lodash.clonedeep";
import { blockchain } from "src/boot/axios";
import { requiredRule } from "src/utils/rules";
import CertificateRequestResult from "./CertificateRequestResult.vue";

const DEFAULT_CERTIFICATE_PARAMS = {
  certificate: "",
  domain: "",
  contract_id: "",
  wallet: "",
};

const useRules = () => {
  return {
    certificate: [requiredRule()],
    domain: [requiredRule()],
  };
};

export default defineComponent({
  name: "RegisterCertificateDialog",
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

const certificateData = ref(cloneDeep(DEFAULT_CERTIFICATE_PARAMS));
const loading = ref(false);

const copy = (value) => {
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

const onSubmit = async () => {
  try {
    loading.value = true;

    const { data } = await blockchain.post("certificate/request", certificateData.value, {
      customErrorHandlers: {
        500: () => {
          $q.notify({
            type: 'negative',
            message: 'The certificate isn\'t valid.'
          })
        }
      }
    });

    Object.assign(certificateData.value, data)

    $q.notify({
      type: 'positive',
      message: data?.message ?? 'Certificate requested sucessfully.'
    })

    $q.dialog({
      component: CertificateRequestResult,
      componentProps: {
        contractId: certificateData.value.contract_id,
        wallet: certificateData.value.wallet
      }
    })
      .onDismiss(() => {
        onOKClick()
      })
  } finally {
    loading.value = false;
  }
};
</script>
