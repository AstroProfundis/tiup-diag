export default [
  {
    path: '/',
    component: '../layouts/BasicLayout',
    Routes: ['src/pages/Authorized'],
    authority: ['admin', 'user'],
    routes: [
      {
        path: '/',
        redirect: '/inspection/instances',
      },
      {
        path: '/inspection',
        name: 'inspection',
        icon: 'dashboard',
        hideChildrenInMenu: true,
        routes: [
          {
            path: '/inspection',
            redirect: '/inspection/instances',
          },
          {
            path: '/inspection/instances',
            name: 'instance_list',
            component: './Inspection/InstanceList',
          },
          {
            path: '/inspection/instances/:id/reports',
            name: 'report_list',
            component: './Inspection/ReportList',
          },
          {
            path: '/inspection/reports/:id',
            name: 'report_detail',
            component: './Inspection/ReportDetail',
          },
        ],
      },
      {
        path: '/misc',
        name: 'misc',
        icon: 'dashboard',
        routes: [
          {
            path: '/misc',
            redirect: '/misc/flamegraphs',
          },
          {
            path: '/misc/flamegraphs',
            name: 'flame_graph',
            component: './Misc/FlameGraph',
          },
          {
            path: '/misc/perfprofiles',
            name: 'perf_profile',
            component: './Misc/PerfProfile',
          },
        ],
      },
      {
        component: './404',
      },
    ],
  },
  {
    component: './404',
  },
];
