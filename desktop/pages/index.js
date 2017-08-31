import Link from 'next/link'
import fetch from 'isomorphic-fetch'

import Layout from '../layouts/application'

const Widget = ({info}) => {
  return (<Layout info={info} user={{}}>
    <div>
      Welcome to next.js!
      <Link href="/admin"><a>Admin</a></Link>
    </div>
  </Layout>)
}

Widget.getInitialProps = async ({ req }) => {
  const res = await fetch('http://localhost:8080/site/info')
  const json = await res.json()
  console.log(process.env.BACKEND)
  return { info: json }
}

export default Widget


//
// export default Widget
