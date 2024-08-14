<template>
  <q-dialog
    ref="dialogRef"
    @hide="onDialogHide"
    @shake="onCloseDialog"
    :position="$q.screen.gt.xs ? 'standard' : 'bottom'"
  >
    <q-card class="medium-dialog">
      <q-form @submit="onSubmit">
        <q-card-section class="row no-wrap">
          <div class="full-width-flex">
            <div class="text-h6 text-grey-9 flex items-center full-height">
              Show Domain
            </div>
          </div>
        </q-card-section>
        <q-separator />
        <q-card-section>
          <div class="row q-col-gutter-md">
            <div class="col-12">
              <q-input
                v-model="hashParams.message"
                label="Message *"
                filled
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="hashParams.key"
                label="Key *"
                filled
                autogrow
              />
            </div>
            <div class="col-12">
              <q-input
                v-model="hash"
                label="Hash"
                autogrow
                readonly
                filled
                input-class="text-caption"
                >
                  <template #append>
                    <q-btn
                      @click="copyKeyToClipboard(hash)"
                      icon="content_copy"
                      size="0.5em"
                      round
                    >
                      <q-tooltip class="text-subtitle2">
                        Copy Hash
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
          <q-btn
            type="submit"
            label="Generate"
            color="primary"
            text-color="white"
            padding="sm md"
            :loading="loading"
          />
        </q-card-actions>
      </q-form>
    </q-card>
  </q-dialog>
</template>

<script>
import { defineComponent, onMounted, ref } from 'vue'
import { useDialogPluginComponent, copyToClipboard, useQuasar } from 'quasar'
import cloneDeep from 'lodash.clonedeep';
import { api } from 'src/boot/axios';

const DEFAULT_HASH_PARAMS = {
  message: '',
  key: '',
}

export default defineComponent({
  name: 'HashDialog'
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

const hash = ref('')
const hashParams = ref(cloneDeep(DEFAULT_HASH_PARAMS))
const loading = ref(false)

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

// TODO: adicionar rules

const onSubmit = async () => {
  hash.value = ''

  try {
    loading.value = true

    const { data } = await api.get('hash', {
      params: hashParams.value,
    })

    hash.value = data.hash

  } finally {
    loading.value = false
  }
}

</script>