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
              Acme Challenge
            </div>
          </div>
        </q-card-section>
        <q-separator />
        <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <q-input
                v-model="acmeChallengeData.contract_id"
                label="Contract ID *"
                :rules="rules.contract_id"
                filled
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="acmeChallengeData.sign"
                label="Signature *"
                :rules="rules.sign"
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
            label="Validate Certificate"
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

const DEFAULT_ACME_CHALLENGE_PARAMS = {
  contract_id: "",
  sign: "",
};

const useRules = () => {
  return {
    contract_id: [requiredRule()],
    sign: [requiredRule()],
  };
};

export default defineComponent({
  name: "AcmeChallengeDialog",
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

const acmeChallengeData = ref(cloneDeep(DEFAULT_ACME_CHALLENGE_PARAMS));
const loading = ref(false);

const onSubmit = async () => {
  try {
    loading.value = true;

    const { data } = await blockchain.post("contract/execute", acmeChallengeData.value, {
      headers: {
        "X-Requested-With": "XMLHttpRequest",
        "Content-Type": "Application/json",
        Authorization: acmeChallengeData.value.sign,
      },
      customErrorHandlers: {
        404: () => {
          $q.notify({
            type: 'negative',
            message: 'Contract not found.'
          })
        },
        401: () => {
          $q.notify({
            type: 'negative',
            message: 'Invalid signature.'
          })
        }
      }
    });

    Object.assign(acmeChallengeData.value, data)

    $q.notify({
      type: 'positive',
      message: data?.message ?? 'Certificate validated successfully.'
    })

    onOKClick()

  } finally {
    loading.value = false;
  }
};
</script>
