import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {Col, Card } from 'antd'

import fail from '../assets/fail.png'
import Layout from './Application'

class Widget extends Component {
  render() {
    const {breadcrumbs, message} = this.props
    return <Layout breadcrumbs={breadcrumbs}>
      <Col md={{offset:8, span:8}}>
        <Card title={message}>
          <img alt="fail" width="100%" src={fail} />
        </Card>
      </Col>
    </Layout>
  }
}


Widget.propTypes = {
  message: PropTypes.node.isRequired,
  breadcrumbs: PropTypes.array.isRequired,
}

export default Widget;
