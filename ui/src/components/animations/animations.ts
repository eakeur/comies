export const ListShowUpAnimation = {
  parent: {
    hidden: { opacity: 1, scale: 0 },
    visible: {
      opacity: 1,
      scale: 1,
      transition: {
        delayChildren: 0.15,
        staggerChildren: 0.05
      }
    }
  },
  children: {
    hidden: { y: 20, opacity: 0 },
    visible: {
      y: 0,
      opacity: 1
    }
  },
}

export const ListAddingAnimation = {
  parent: {
    hidden: { x: 100 },
    visible: { 
      x: 0, 
      transition: { 
        type: "spring", 
        stiffness: 100 
      } 
    }
  },
}