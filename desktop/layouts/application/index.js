import React from 'react'
import PropTypes from 'prop-types'

import Header from './Header'
import Footer from './Footer'

const Widget = ({children, info}) => (
  <div>
    <Header info={info}/>
    {children}
    <Footer info={info}/>
  </div>
)


Widget.propTypes = {
  children: PropTypes.node.isRequired,
  user: PropTypes.object.isRequired,
  info: PropTypes.object.isRequired,
}

export default Widget
