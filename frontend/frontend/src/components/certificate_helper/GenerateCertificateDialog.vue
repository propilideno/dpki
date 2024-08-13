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
            Generate Certificate
          </div>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section>
        <div class="row q-col-gutter-md">
          <div
            v-for="key, index in Object.keys(keys)"
            :key="index"
            class="col-12"
          >
            <q-input
              v-model="keys[key]"
              :label="formatKeyLabel(key)"
              autogrow
              readonly
              outlined
              filled
              input-class="text-caption"
            >
              <template #append>
                <q-btn
                  @click="copyKeyToClipboard(keys[key])"
                  icon="content_copy"
                  size="0.5em"
                  round
                >
                  <q-tooltip class="text-subtitle2">
                    Copy {{ formatKeyLabel(key) }}
                  </q-tooltip>
                </q-btn>
              </template>
            </q-input>
          </div>
          <div class="col-12"></div>
          <div class="col-12"></div>
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
          @click="generateCertificate"
          label="Generate"
          color="primary"
          text-color="white"
          padding="sm md"
          :loading="loadingKeys"
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue'
import { useDialogPluginComponent, copyToClipboard, useQuasar } from 'quasar'
import cloneDeep from 'lodash.clonedeep';
import { api } from 'src/boot/axios';

const DEFAULT_KEYS = {
  private_key: '',
  public_key: '',
  certificate: '',
}

export default defineComponent({
  name: 'GenerateCertificateDialog'
})
</script>

<script setup>

/* const props = */ defineProps({})

/* const emit = */ defineEmits([...useDialogPluginComponent.emits])

const { dialogRef, onDialogHide, onDialogOK: onOKClick, onDialogCancel: onCancelClick } = useDialogPluginComponent()

const onCloseDialog = () => {
  onCancelClick()
}

const $q = useQuasar()

const keys = ref(cloneDeep(DEFAULT_KEYS))
const loadingKeys = ref(false)

const copyKeyToClipboard = (value) => {
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

const formatKeyLabel = (key) => {
  const labels = {
    private_key: 'Private Key',
    public_key: 'Public Key',
    certificate: 'Certificate',
  }

  return labels[key]
}

const generateCertificate = async () => {
  try {
    loadingKeys.value = true

    const { data } = await api.get('certificate/new')

    Object.assign(keys.value, data)

  } finally {
    loadingKeys.value = false
  }
}

onMounted(async () => {
  await generateCertificate()
})

</script>