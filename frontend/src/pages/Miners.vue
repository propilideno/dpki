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
import NewCertificateDialog from 'src/components/certificate_manager/RegisterCertificateDialog.vue'

const $q = useQuasar()

const myFunction = () => {
  $q.dialog({
    component: NewCertificateDialog
  })
}

const actions = [
  {
    label: 'Mine Block',
    icon: 'grade',
    onClick: myFunction
  },
  {
    label: 'Mine Transactions',
    icon: 'attach_money',
    onClick: myFunction
  },
  {
    label: 'Mine Contract',
    icon: 'receipt_long',
    onClick: myFunction
  },
]

</script>
