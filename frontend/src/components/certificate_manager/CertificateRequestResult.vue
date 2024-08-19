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
            Data from Certificate requested
          </div>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <div class="row q-col-gutter-md">
          <div class="col-12">
            <q-input
              v-model="innerContractId"
              label="Contract ID"
              readonly
              filled
              autogrow
              >
                <template #append>
                  <q-btn
                    @click="copy(innerContractId)"
                    icon="content_copy"
                    size="0.5em"
                    round
                  >
                    <q-tooltip class="text-subtitle2">
                      Copy Contract ID
                    </q-tooltip>
                  </q-btn>
                </template>
              </q-input>
          </div>
          <div class="col-12">
            <q-input
              v-model="innerWallet"
              label="Wallet"
              readonly
              filled
              autogrow
              >
                <template #append>
                  <q-btn
                    @click="copy(innerWallet)"
                    icon="content_copy"
                    size="0.5em"
                    round
                  >
                    <q-tooltip class="text-subtitle2">
                      Copy Wallet
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
  contractId: {
    type: String,
    required: true
  },
  wallet: {
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

const innerContractId = ref(props.contractId)
const innerWallet = ref(props.wallet)

const copy = (value) => {
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