<template>
  <q-page padding>
    <q-form ref="formRef">
      <div class="row q-col-gutter-lg">
        <div class="col-6">
          <q-list bordered class="rounded-borders q-pa-none q-mb-md">
            <q-expansion-item
              icon="swap_vert"
              label="Transaction Pool"
              header-class="text-subtitle1"
              default-opened
            >
              <q-card>
                <q-card-section>
                  <q-list separator>
                    <q-item>
                      <q-item-section>
                        <q-item-label>
                          From
                        </q-item-label>
                      </q-item-section>
                      <q-item-section>
                        <q-item-label>
                          To
                        </q-item-label>
                      </q-item-section>
                      <q-item-section>
                        <q-item-label class="text-center">
                          Amount
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                    <q-item v-if="memoryPoolLoading">
                      <q-item-section>
                        <q-item-label class="text-center">
                          <q-spinner
                            color="primary"
                            size="2em"
                          />
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                    <q-item v-else-if="memoryPool.transactionpool.length === 0">
                      <q-item-section>
                        <q-item-label caption>
                          No transactions.
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                    <q-item
                      v-else
                      v-for="item, index in memoryPool.transactionpool"
                      :key="index"
                      clickable
                    >
                      <q-item-section>
                        <q-item-label lines="1">
                          {{ item.from }}
                        </q-item-label>
                      </q-item-section>
                      <q-item-section>
                        <q-item-label lines="1">
                          {{ item.to }}
                        </q-item-label>
                      </q-item-section>
                      <q-item-section>
                        <q-item-label class="text-center">
                          {{ item.amount }}
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                  </q-list>
                </q-card-section>
              </q-card>
            </q-expansion-item>
          </q-list>
          <q-list bordered class="rounded-borders q-pa-none">
            <q-expansion-item
              icon="swap_vert"
              label="Contract Execution Pool"
              header-class="text-subtitle1"
              default-opened
            >
              <q-card>
                <q-card-section>
                  <q-list separator>
                    <q-item>
                      <q-item-section>
                        <q-item-label>
                          Timestamp
                        </q-item-label>
                      </q-item-section>
                      <q-item-section>
                        <q-item-label>
                          Contract ID
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                    <q-item v-if="memoryPoolLoading">
                      <q-item-section>
                        <q-item-label class="text-center">
                          <q-spinner
                            color="primary"
                            size="2em"
                          />
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                    <q-item v-else-if="memoryPool.contractexecutionpool.length === 0">
                      <q-item-section>
                        <q-item-label caption>
                          No contract executions.
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                    <q-item
                      v-else
                      v-for="item, index in memoryPool.contractexecutionpool"
                      :key="index"
                      clickable
                    >
                      <q-item-section>
                        <q-item-label lines="1">
                          {{ formatDate(item.timestamp) }}
                        </q-item-label>
                      </q-item-section>
                      <q-item-section>
                        <q-item-label lines="1">
                          {{ item.contract_id }}
                        </q-item-label>
                      </q-item-section>
                    </q-item>
                  </q-list>
                </q-card-section>
              </q-card>
            </q-expansion-item>
          </q-list>
        </div>
        <div class="col-6">
          <div class="row q-col-gutter-md">
            <div class="col-12">
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
              :class="action?.class ?? 'col-6'"
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
import { onMounted, ref } from 'vue'
import { date, useQuasar } from 'quasar'
import { requiredRule } from 'src/utils/rules';
import { blockchain } from 'src/boot/axios';
import cloneDeep from 'lodash.clonedeep';

const DEFAULT_MEMORY_POOL = {
  contractexecutionpool: [],
  transactionpool: [],
}

export default {
  name: 'DPKICertificate',
}
</script>

<script setup>

const $q = useQuasar()

const formRef = ref(null)

const loadingMine = ref(false)
const mineBlockchain = async (url, message) => {
  if (!await formRef.value?.validate()) return

  try {
    loadingMine.value = true

    await blockchain.post(url, { wallet: wallet.value })

    $q.notify({
      type: 'positive',
      message
    })

    await getMemoryPool()
  } finally {
    loadingMine.value = false
  }
}

const actions = [
  {
    label: 'Mine Block',
    icon: 'grade',
    onClick: () => mineBlockchain('mine/block', 'Block mined sucessfully.')
  },
  {
    label: 'Mine Transactions',
    icon: 'attach_money',
    onClick: () => mineBlockchain('mine/transaction', 'Transaction mined sucessfully.')
  },
  {
    label: 'Mine Contract Execution',
    icon: 'receipt_long',
    onClick: () => mineBlockchain('mine/contract', 'Contract execution mined sucessfully.'),
    class: 'col-12'
  },
]

const wallet = ref('')

const memoryPool = ref(cloneDeep(DEFAULT_MEMORY_POOL))

const memoryPoolLoading = ref(false)
const getMemoryPool = async () => {
  memoryPool.value = cloneDeep(DEFAULT_MEMORY_POOL)

  try {
    memoryPoolLoading.value = true

    const { data } = await blockchain.get('memorypool')

    memoryPool.value.contractexecutionpool = data?.contractexecutionpool ?? []
    memoryPool.value.transactionpool = data?.transactionpool ?? []
  } finally {
    memoryPoolLoading.value = false
  }
}

const formatDate = (timestamp) => {
  return date.formatDate(timestamp, 'YYYY-MM-DD HH:mm:ss')
}

onMounted(async () => {
  await getMemoryPool()
})

</script>
