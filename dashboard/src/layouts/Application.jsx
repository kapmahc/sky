import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import {Layout, Menu, Row, Breadcrumb} from 'antd'
import {FormattedMessage} from 'react-intl'
import {Link} from 'react-router-dom'

import Root from './Root'
import Footer from '../components/Footer'

const { Header, Content } = Layout

class Widget extends Component {
  render() {
    const {children, info, breadcrumbs} = this.props
    return (<Root>
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
          <Breadcrumb.Item><Link to="/"><FormattedMessage id="site.home.title"/></Link></Breadcrumb.Item>
          {breadcrumbs.map((l,i)=> <Breadcrumb.Item key={i}><Link to={l.href}>{l.label}</Link></Breadcrumb.Item>)}
        </Breadcrumb>
        <Row style={{ background: '#fff', padding: 24, minHeight: 380 }}>
          {children}
        </Row>
      </Content>
      <Footer/>
    </Root>)
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
