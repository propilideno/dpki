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
import CertificateStatusDialog from 'src/components/certificate_manager/CertificateStatusDialog.vue';
import RegisterCertificateDialog from 'src/components/certificate_manager/RegisterCertificateDialog.vue'

export default {
  name: 'DPKICertificateManager',
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
    label: 'Request Certificate',
    icon: 'verified_user',
    onClick: () => openDialog(RegisterCertificateDialog)
  },
  {
    label: 'Check Acme Challenge Situation',
    icon: 'emoji_events',
    onClick: () => openDialog(RegisterCertificateDialog)
  },
  {
    label: 'Certificate Status',
    icon: 'query_stats',
    onClick: () => openDialog(CertificateStatusDialog)
  },
]

</script>
