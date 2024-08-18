<template>
  <q-page padding>
    <q-form ref="formRef">
      <div class="row q-col-gutter-lg full-width">
        <div class="col-6">
          Listagem
        </div>
        <div class="col-6">
          <div class="column q-col-gutter-md">
            <div class="col-4">
              <q-input
                label="Wallet *"
                v-model="wallet"
                :rules="[requiredRule()]"
                filled
              />
            </div>
            <div
              v-for="action, index in actions"
              :key="index"
              class="col-4"
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
        </div>
      </div>
    </q-form>
  </q-page>
</template>

<script>
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/rules';
import { blockchain } from 'src/boot/axios';

export default {
  name: 'DPKICertificate',
}
</script>

<script setup>

const $q = useQuasar()

const formRef = ref(null)

const loadingMine = ref(false)
const mineBlockchain = async (url) => {
  if (!await formRef.value?.validate()) return

  try {
    loadingMine.value = true

    await blockchain.post(url, { wallet: wallet.value })

    $q.notify({
      type: 'positive',
      message: 'Block mined sucessfully!'
    })
  } finally {
    loadingMine.value = false
  }
}

const actions = [
  {
    label: 'Mine Block',
    icon: 'grade',
    onClick: () => mineBlockchain('mine/block')
  },
  {
    label: 'Mine Transactions',
    icon: 'attach_money',
    onClick: () => mineBlockchain('mine/transaction')
  },
  {
    label: 'Mine Contract',
    icon: 'receipt_long',
    onClick: () => mineBlockchain('mine/contract')
  },
]

const wallet = ref('')

</script>
