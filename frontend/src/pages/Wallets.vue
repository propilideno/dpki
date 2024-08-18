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

export default {
  name: 'DPKICertificate',
}
</script>

<script setup>
import RegisterCertificateDialog from 'src/components/certificate_manager/RegisterCertificateDialog.vue'

const $q = useQuasar()

const myFunction = () => {
  $q.dialog({
    component: RegisterCertificateDialog
  })
}

const actions = [
  {
    label: 'Show Wallet',
    icon: 'credit_card',
    onClick: myFunction
  },
  {
    label: 'Transfer Money',
    icon: 'paid',
    onClick: myFunction
  },
  {
    label: 'Get Wallets',
    icon: 'groups',
    onClick: myFunction
  },
]

</script>
