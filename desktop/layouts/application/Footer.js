import React from 'react'
import PropTypes from 'prop-types'

const Widget = ({info}) => (
  <div>
    application footer
  </div>
)

Widget.propTypes = {
  info: PropTypes.object.isRequired,
}

export default Widget
