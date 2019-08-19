export default [
  {
    path: '/user',
    component: '../layouts/UserLayout',
    routes: [
      { path: '/user', redirect: '/user/login' },
      { path: '/user/login', name: 'login', component: './User/Login' },
      {
        component: '404',
      },
    ],
  },
  {
    path: '/',
    component: '../layouts/BasicLayout',
    routes: [
      {
        path: '/',
        redirect: '/inspection',
      },
      {
        path: '/inspection',
        name: 'inspection',
        icon: 'dashboard',
        hideChildrenInMenu: true,
        routes: [
          {
            path: '/inspection',
            name: 'inspection',
            component: './Inspection/InspectionHome',
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
            path: '/inspection/instances/:instanceId/reports/:id',
            name: 'report_detail',
            component: './Inspection/ReportDetail',
          },
          {
            path: '/inspection/reports/:id',
            name: 'report_detail',
            component: './Inspection/ReportDetail',
          },
          {
            path: '/inspection/reports',
            name: 'report_list',
            component: './Inspection/ReportList',
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
            component: './Misc/FlameGraphList',
          },
          {
            path: '/misc/flamegraphs/:id',
            name: 'flame_graph_detail',
            component: './Misc/FlameGraphDetail',
            hideInMenu: true,
          },
          {
            path: '/misc/perfprofiles',
            name: 'perf_profile',
            component: './Misc/PerfProfileList',
          },
          {
            path: '/misc/perfprofiles/:id',
            name: 'perf_profile_detail',
            component: './Misc/PerfProfileDetail',
            hideInMenu: true,
          },
        ],
      },
      {
        path: '/logs',
        name: 'logs',
        icon: 'dashboard',
        component: './Log/LogList',
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