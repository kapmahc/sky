import React, { Component } from 'react'

import Layout from '../../layouts/Application'

class Widget extends Component {
  render() {
    return <Layout breadcrumbs={[]}>
      <h1>home</h1>
    </Layout>
  }
}

export default Widget;
