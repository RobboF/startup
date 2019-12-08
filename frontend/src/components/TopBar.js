import React from 'react'
import { Button, Pane, Heading, defaultTheme } from "evergreen-ui";

export default function TopBar() {
  return (
    <Pane display="flex" padding={16} background={defaultTheme.palette.red.base} borderRadius={3}>
    <Pane flex={1} alignItems="center" display="flex" justifyContent="center">
      <Heading size={700}  color="white">Yggdrasil</Heading>
    </Pane>
  </Pane>
  )
}
