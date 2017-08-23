import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import {Layout, Menu, Row, Breadcrumb} from 'antd'

import Application from './Application'
import Footer from '../components/Footer'

const { Header, Content } = Layout

class Widget extends Component {
  render() {
    const {children, info, breadcrumbs} = this.props
    return (<Application>
      <Header>
        <div className="logo" />
        <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={[]}
          style={{ lineHeight: '64px' }}
        >
          <Menu.Item key="home">{info.title}</Menu.Item>
        </Menu>
      </Header>
      <Content style={{ padding: '0 50px'}}>
        <Breadcrumb style={{ margin: '12px 0' }}>
          {breadcrumbs.map((b,i)=> <Breadcrumb.Item key={i}>{b.label}</Breadcrumb.Item>)}
        </Breadcrumb>
        <Row style={{ background: '#fff', padding: 24, minHeight: 380 }}>
          {children}
        </Row>
      </Content>
      <Footer/>
    </Application>)
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  info: PropTypes.object.isRequired,
  breadcrumbs: PropTypes.array.isRequired,
}

export default connect(
  state => ({
    info: state.siteInfo,
  }),
  {},
)(Widget)
