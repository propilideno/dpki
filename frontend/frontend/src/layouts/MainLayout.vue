<template>
  <q-layout view="hHh LpR fFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          Descentralized Public Key Infrastructure
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-tabs
          indicator-color="transparent"
          active-class="active-item"
          vertical
        >
          <q-route-tab
            v-for="link in linksList"
            :key="link.title"
            :label="link.title"
            v-bind="link"
          >
            <!-- <q-item>
              <q-item-section side>
                <q-icon :name="link.icon" size="md" />
              </q-item-section>
              <q-item-section>
                <q-item-label>
                  {{ link.label }}
                </q-item-label>
              </q-item-section>
            </q-item> -->
          </q-route-tab>
      </q-tabs>
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from 'vue'

defineOptions({
  name: 'MainLayout'
})

const linksList = [
  {
    title: 'Certificate Helper',
    caption: 'Generate a self signed certificate and a public/private key pair.',
    icon: 'key',
    to: { name: 'DPKI.CertificateHelper' }
  },
  {
    title: 'Wallets',
    caption: 'Prove that you are yourself to valid your domain.',
    icon: 'wallet',
    to: { name: 'DPKI.Wallets' }
  },
  {
    title: 'Miners',
    caption: 'Mine in Blockchain and get recompenses.',
    icon: 'construction',
    to: { name: 'DPKI.Miners' }
  },
  {
    title: 'DNS',
    caption: 'Prove that you are yourself to valid your domain.',
    icon: 'dns',
    to: { name: 'DPKI.Dns' }
  },
  {
    title: 'Certificate Manager',
    caption: 'Register, revoke and consult certificates in Blockchain.',
    icon: 'badge',
    to: { name: 'DPKI.CertificateManager' }
  },
]

const leftDrawerOpen = ref(false)

function toggleLeftDrawer () {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>
