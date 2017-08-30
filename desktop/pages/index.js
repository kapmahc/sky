import Link from 'next/link'

import Layout from '../layouts/application'

export default() => (
  <Layout>
    <div>
      Welcome to next.js!
      <Link href="/admin"><a>Admin</a></Link>
    </div>
  </Layout>
)
