<template>
  <q-dialog
    ref="dialogRef"
    @hide="onDialogHide"
    @shake="onCloseDialog"
    :position="$q.screen.gt.xs ? 'standard' : 'bottom'"
  >
    <q-card class="full-width">
      <q-card-section class="row no-wrap">
        <div class="full-width-flex">
          <div class="text-h6 text-grey-9 flex items-center full-height text-roboto-slab">
            Criar novo certificado
          </div>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <!-- TODO -->
      </q-card-section>
      <q-card-actions align="right">
        <q-btn
          @click="onCloseDialog"
          label="Fechar"
          color="grey-9"
          padding="sm md"
          flat
        />
        <q-btn
          @click="onOKClick"
          label="Continuar"
          color="primary"
          text-color="white"
          padding="sm md"
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent } from 'vue'
import { useDialogPluginComponent } from 'quasar'

export default defineComponent({
  name: 'NewCertificateDialog'
})
</script>

<script setup>

/* const props = */ defineProps({})

/* const emit = */ defineEmits([...useDialogPluginComponent.emits])

const { dialogRef, onDialogHide, onDialogOK: onOKClick, onDialogCancel: onCancelClick } = useDialogPluginComponent()

const onCloseDialog = () => {
  onCancelClick()
}

</script>