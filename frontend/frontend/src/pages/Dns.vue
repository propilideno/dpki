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
  name: 'DPKIDns',
}
</script>

<script setup>
import RegisterDomainDialog from 'src/components/dns/RegisterDomainDialog.vue';
import DeleteRegisterDomainDialog from 'src/components/dns/DeleteRegisterDomainDialog.vue';
import ShowRegisterDomainDialog from 'src/components/dns/ShowRegisterDomainDialog.vue';

const $q = useQuasar()

const openDialog = (component) => {
  $q.dialog({
    component: component
  })
}

const actions = [
  {
    label: 'Register Domain',
    icon: 'domain_add',
    onClick: () => openDialog(RegisterDomainDialog)
  },
  {
    label: 'Delete Register Domain',
    icon: 'domain_disabled',
    onClick: () => openDialog(DeleteRegisterDomainDialog)
  },
  {
    label: 'Show Register Domain',
    icon: 'domain',
    onClick: () => openDialog(ShowRegisterDomainDialog)
  },
]

</script>
