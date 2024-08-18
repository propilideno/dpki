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
            Shared Key
          </div>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <p>
          Your shared key is:
        </p>
        <div class="row q-col-gutter-md">
          <div class="col-12">
            <q-input
              v-model="innerSharedKey"
              label="Shared Key"
              readonly
              filled
              autogrow
              >
                <template #append>
                  <q-btn
                    @click="copySharedKeyToClipboard(innerSharedKey)"
                    icon="content_copy"
                    size="0.5em"
                    round
                  >
                    <q-tooltip class="text-subtitle2">
                      Copy Shared Key
                    </q-tooltip>
                  </q-btn>
                </template>
              </q-input>
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
import { defineComponent, ref } from 'vue'
import { useDialogPluginComponent, copyToClipboard, useQuasar } from 'quasar'

export default defineComponent({
  name: 'HashDialog'
})
</script>

<script setup>

const props = defineProps({
  sharedKey: {
    type: String,
    required: true
  }
})

/* const emit = */ defineEmits([...useDialogPluginComponent.emits])

const { dialogRef, onDialogHide, onDialogOK: onOKClick, onDialogCancel: onCancelClick } = useDialogPluginComponent()

const onCloseDialog = () => {
  onCancelClick()
}

const $q = useQuasar()

const innerSharedKey = ref(props.sharedKey)

const copySharedKeyToClipboard = (value) => {
  copyToClipboard(value)
    .then(() => {
      $q.notify({
        type: 'positive',
        message: 'Copied!'
      })
    })
    .catch(() => {
      $q.notify({
        type: 'alert',
        message: 'Error on Copy.'
      })
    })
}

</script>