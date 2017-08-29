export const TOKEN = 'token'

export const DATE_FORMAT = 'YYYY/MM/DD'

export const dashboard = (user) => {
  var items = []
  // is sign-in ?
  if (user.uid) {
    items.push({
      icon: 'user',
      label: 'auth.dashboard.title',
      items: [
        {to: '/users/logs', label: 'auth.users.logs.title'},
        {to: '/users/change-password', label: 'auth.users.change-password.title'},
        {to: '/users/info', label: 'auth.users.info.title'},
        {to: '/attachments', label: 'auth.attachments.index.title'},
      ],
    })
    // is-admin?
    if(user.admin){
      items.push({
        label: 'site.dashboard.title',
        icon: 'setting',
        items: [
          {to: '/admin/status', label: 'site.admin.status.title'},
          {to: '/admin/site/info', label: 'site.admin.site.info.title'},
          {to: '/admin/site/author', label: 'site.admin.site.author.title'},
          {to: '/admin/seo', label: 'site.admin.seo.title'},
          {to: '/admin/smtp', label: 'site.admin.smtp.title'},
          {to: '/admin/paypal', label: 'site.admin.paypal.title'},
          {to: '/admin/locales', label: 'site.admin.locales.index.title'},
          {to: '/admin/links', label: 'site.admin.links.index.title'},
          {to: '/admin/cards', label: 'site.admin.cards.index.title'},
          {to: '/admin/friend-links', label: 'site.admin.friend-links.index.title'},
          {to: '/admin/users', label: 'site.admin.users.index.title'},
          {to: '/leave-words', label: 'site.leave-words.index.title'},
        ]
      })

      items.push({
        icon: 'question-circle-o',
        label: 'survey.dashboard.title',
        items: [
          {to:'/survey/forms', label: 'survey.forms.index.title'}
        ],
      })

    }

  }
  return items
}
