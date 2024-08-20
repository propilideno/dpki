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
              Transfer Money
            </div>
          </div>
        </q-card-section>
        <q-separator />
        <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <q-input
                v-model="transferData.from"
                :rules="rules.from"
                label="From Wallet *"
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="transferData.to"
                :rules="rules.to"
                label="To Wallet *"
              />
            </div>
            <div class="col-12">
              <q-input
                v-model.number="transferData.amount"
                label="Amount *"
                :rules="rules.amount"
                type="number"
                inputmode="numeric"
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
            label="Transfer"
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
import { blockchain } from "src/boot/axios";
import { requiredRule } from "src/utils/rules";

const DEFAULT_TRANSFER = {
  from: "",
  to: "",
  amount: null,
};

const useRules = () => {
  return {
    from: [requiredRule()],
    to: [requiredRule()],
    amount: [requiredRule()],
  };
};

export default defineComponent({
  name: "TransferMoneyDialog",
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

const transferData = ref(cloneDeep(DEFAULT_TRANSFER));

const loading = ref(false);
const onSubmit = async () => {
  try {
    loading.value = true;

    await blockchain.post('transaction/new', transferData.value);

    $q.notify({
      type: 'positive',
      message: 'Transfered successfully.'
    })

    onOKClick()
  } finally {
    loading.value = false;
  }
};

</script>
