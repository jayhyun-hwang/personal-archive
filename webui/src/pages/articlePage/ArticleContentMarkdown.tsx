import React, { FC } from "react"
import { Button } from "rsuite"
import { useHistory } from "react-router-dom"
import { If } from "react-if"
import MarkdownContent from "../../component/common/MarkdownContent"
import Article from "../../models/Article"
import GithubRepoWidget from "../../component/GithubRepoWidget"


interface Props {
  article: Article
}

const ArticleContentMarkdown: FC<Props> = ({article}) => {
  const history = useHistory()

  return (
    <>
      <div style={{textAlign: 'right'}}>
        <Button
          appearance="link"
          onClick={() => history.push(`/articles/${article.id}/edit`)}
        >
          EDIT
        </Button>
      </div>
      <If condition={isGithubRepoURL(article.url)}>
        <GithubRepoWidget
          user={getGithubUser(article.url)}
          repo={getGithubRepo(article.url)}
        />
      </If>
      <MarkdownContent content={article.content}/>
    </>
  )
}

const githubRepoRegex = /^https:\/\/github\.com\/([a-zA-Z0-9\-]*)\/([a-zA-Z0-9\-]*)/gi
const isGithubRepoURL = (url: string): boolean => githubRepoRegex.test(url)
const getGithubUser = (url: string): string => githubRepoRegex.exec(url)![1]
const getGithubRepo = (url: string): string => githubRepoRegex.exec(url)![2]

export default ArticleContentMarkdown
