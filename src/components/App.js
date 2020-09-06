import React from 'react'
import { SWRConfig } from 'swr'
import axios from 'axios'

import './App.css'

import Posts from './Posts'

export default function App () {

  return (
    <SWRConfig
      value={{
        fetcher: url => axios.get(url).then(res => res.data)
      }}
    >
      <Posts />
    </SWRConfig>
  )
}
