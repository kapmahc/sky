import Document, { Head, Main, NextScript } from 'next/document'
import flush from 'styled-jsx/server'

class Widget extends Document {
  static async getInitialProps (context) {
    const props = await super.getInitialProps(context)
    // const {req: {localeDataScript}} = context
    return {
      ...props,
      // localeDataScript
    }
  }

  render () {
    return (
     <html>
       <Head />
       <body>
         <Main />
         <NextScript />
       </body>
     </html>
    )
  }
}

export default Widget
