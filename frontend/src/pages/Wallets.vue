<template>
  <q-page class="flex flex-center" padding>
    <div class="row q-col-gutter-lg full-width">
      <div
        v-for="action, index in actions"
        :key="index"
        class="col"
      >
        <q-btn
          @click="action.onClick"
          :label="action.label"
          :icon="action.icon"
          stack
          size="3em"
          padding="xl"
          style="border-radius: 16px;"
          no-caps
          class="fit"
          color="grey-1"
          text-color="black"
        />
      </div>
    </div>
  </q-page>
</template>

<script>
import { useQuasar } from 'quasar'
import ShowWalletDialog from 'src/components/wallets/ShowWalletDialog.vue';
import TransferMoneyDialog from 'src/components/wallets/TransferMoneyDialog.vue';

export default {
  name: 'DPKICertificate',
}
</script>

<script setup>

const $q = useQuasar()

const openDialog = (component) => {
  $q.dialog({
    component
  })
}

const actions = [
  {
    label: 'Show Wallet',
    icon: 'credit_card',
    onClick: () => openDialog(ShowWalletDialog)
  },
  {
    label: 'Transfer Money',
    icon: 'paid',
    onClick: () => openDialog(TransferMoneyDialog)
  },
]

</script>
