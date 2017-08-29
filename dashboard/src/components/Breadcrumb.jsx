import React from 'react'
import PropTypes from 'prop-types'
import {FormattedMessage} from 'react-intl'
import {Breadcrumb} from 'antd'
import {Link} from 'react-router-dom'

const Widget = ({items}) => <Breadcrumb style={{ margin: '12px 0' }}>
  <Breadcrumb.Item><Link to="/"><FormattedMessage id="breadcrumb.home"/></Link></Breadcrumb.Item>
  {items.map((l,i)=> <Breadcrumb.Item key={i}><Link to={l.href}>{l.label}</Link></Breadcrumb.Item>)}
</Breadcrumb>

Widget.propTypes = {
  items: PropTypes.array.isRequired,
}

export default Widget
