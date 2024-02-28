import { useEffect, useState } from 'react';
import axios from 'axios'

export default function Home() {
  
  useEffect(() => {
    document.title = "Home Page"
  }, [])

  return (
      <div className="container-xl">
        test
      </div>
  )
}