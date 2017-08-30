import Link from 'next/link'

import Layout from '../../layouts/dashboard'

export default() => (
  <Layout>
    <div>
      Welcome to next.js!
      <Link href="/"><a>Home</a></Link>
    </div>
  </Layout>
)
