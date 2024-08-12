const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        component: () => import('pages/Home.vue')
      },
      {
        path: '/certificate-helper',
        name: 'DPKI.CertificateHelper',
        component: () => import('pages/CertificateHelper.vue')
      },{
        path: '/wallets',
        name: 'DPKI.Wallets',
        component: () => import('pages/Wallets.vue')
      },
      {
        path: '/miners',
        name: 'DPKI.Miners',
        component: () => import('pages/Miners.vue')
      },
      {
        path: '/dns',
        name: 'DPKI.Dns',
        component: () => import('pages/Dns.vue')
      },
      {
        path: '/certificate-manager',
        name: 'DPKI.CertificateManager',
        component: () => import('pages/CertificateManager.vue')
      },
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
