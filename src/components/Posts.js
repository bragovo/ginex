import React from 'react'
import useSWR from 'swr'

import styles from './Posts.module.css'

export default function Posts () {
  const { data, error } = useSWR('/api/posts')

  return (
    <div className={styles.root}>
      <h1>
        Posts
      </h1>

      {data &&
        <div>
          {data.posts.map(post =>
            <div key={post.Id}>
              {post.Title}
            </div>
          )}
        </div>
      }
    </div>
  )
}
