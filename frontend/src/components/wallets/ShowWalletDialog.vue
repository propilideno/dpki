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
            Show Wallet
          </div>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <div class="row q-col-gutter-md">
          <div class="col-12">
            <q-input
              @update:model-value="onUpdateWallet"
              :model-value="walletData.wallet"
              :rules="rules.wallet"
              label="Wallet *"
              debounce="1000"
              :loading="walletLoading"
            />
          </div>
          <div v-if="walletData.exists" class="col-12">
            <q-input
              v-model="walletData.balance"
              label="Balance"
              readonly
              autogrow
              filled
              stack-label
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
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, ref, reactive } from "vue";
import { useDialogPluginComponent, useQuasar } from "quasar";
import cloneDeep from "lodash.clonedeep";
import { blockchain } from "src/boot/axios";
import { requiredRule } from "src/utils/rules";

const DEFAULT_WALLET_DATA = {
  wallet: "",
  balance: null,
  exists: false,
};

const useRules = () => {
  return {
    wallet: [requiredRule()],
  };
};

export default defineComponent({
  name: "ShowWalletDialog",
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

const walletData = ref(cloneDeep(DEFAULT_WALLET_DATA));

const walletLoading = ref(false);
const walletLoaded = ref(false);
const getWallet = async () => {
  walletData.value.exists = false;
  walletLoaded.value = false;

  try {
    walletLoading.value = true;

    const { data } = await blockchain.get(`info/${walletData.value.wallet}`);

    $q.notify({
      type: 'positive',
      message: 'Balance sucessfully obtained.'
    })

    walletData.value.exists = true;

    Object.assign(walletData.value, data);
  } finally {
    walletLoading.value = false;
    walletLoaded.value = true;
  }
};

const onUpdateWallet = async (newWallet) => {
  walletData.value.wallet = newWallet;

  await getWallet();
};

</script>
