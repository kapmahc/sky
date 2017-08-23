import React, {Component} from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import { Layout } from 'antd'
import {FormattedMessage} from 'react-intl'

import {setLocale} from '../intl'

const {Footer } = Layout;

class Widget extends Component {
  onSwitchLang (l) {
    console.log(l)
  }
  render () {
    const {info} = this.props
    return <Footer style={{ textAlign: 'center' }}>
      &copy; {info.copyright}
      &middot; <FormattedMessage id="footer.others"/> {info.languages.map((l, i)=><a onClick={() => setLocale(l)} style={{ padding: '0 2px'}} key={i}><FormattedMessage id={`languages.${l}`}/></a>)}
    </Footer>
  }
}
Widget.propTypes = {
  info: PropTypes.object.isRequired,
}

export default connect(
  state => ({
    info: state.siteInfo,
  }),
  {},
)(Widget)
