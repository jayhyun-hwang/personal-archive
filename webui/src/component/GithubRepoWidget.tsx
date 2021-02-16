import React, { FC, useEffect, useState } from "react"
import { Alert, Loader, Panel } from "rsuite"


interface Props {
  user: string
  repo: string
}

const GithubRepoWidget: FC<Props> = ({ user: userInput, repo: repoInput }) => {
  const [ isFetching, setFetching ] = useState(false)
  const [ user, setUser ] = useState(emptyUser)
  const [ repo, setRepo ] = useState(emptyRepo)

  console.log({ userInput, repoInput })

  useEffect(() => {
    setFetching(true)
    window.fetch(`https://api.github.com/repos/${userInput}/${repoInput}`)
      .then(resp => resp.json())
      .then(resp => {
        const {
          owner,
          full_name: fullName,
          language = '',
          description = '',
          stargazers_count: starCount = 0,
          forks: forkCount = 0,
        } = resp
        const {
          login,
          avatar_url: avatarURL,
        } = owner

        setUser({
          login,
          avatarURL,
        })
        setRepo({
          fullName,
          language,
          description,
          starCount,
          forkCount,
        })
      })
      .catch(err => Alert.error(err.toString()))
      .finally(() => setFetching(false))
  }, [])

  if (isFetching) {
    return <Loader />
  }

  return (
    <Panel shaded>
      <img alt={user.login} src={user.avatarURL} />
      <div>asdf</div>
    </Panel>
  )
  // TODO IMME
}

interface User {
  login: string
  avatarURL: string
}

interface Repo {
  fullName: string
  language: string
  description: string
  starCount: number
  forkCount: number
}

const emptyUser: User = {
  login: '',
  avatarURL: '',
}

const emptyRepo: Repo = {
  fullName: '',
  language: '',
  description: '',
  starCount: -1,
  forkCount: -1,
}

export default GithubRepoWidget
