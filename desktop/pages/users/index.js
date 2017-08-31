import React from 'react'
import Error from 'next/error'
// import fetch from 'isomorphic-fetch'

export default class Page extends React.Component {
  componentDidMount () {
    console.log('aaa')
  }
  static async getInitialProps () {
    const res = await fetch('http://localhost:8080/site/info')
    const statusCode = res.statusCode > 200 ? res.statusCode : false
    const json = await res.json()
    console.log('bbb')
    return { statusCode, stars: json.languages.length }
  }

  render () {
    if(this.props.statusCode) {
        return <Error statusCode={this.props.statusCode} />
    }

    return (
      <div>Next stars: {this.props.stars}</div>
    )
  }
}
