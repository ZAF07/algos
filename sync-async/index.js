// const prom = new Promise((resolve, reject) => {
//   setTimeout(()=> {
//     reject('Third')
//   }, 2000)
// })

//  Using the callback method, we dont have to block the main function/thread. All tasks are executed in order and once the longer executing function is done, it invokes the callback fn which simply logs the results to console
const cb = (s) => {
    console.log("result here in callback: ", s);
}

const asyncFn = (callback) => {
  const result = new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve("DONE DEAL!")
    }, 2000)
  }) 
// return result
  result.then(res => {
    callback(res)
  })
}

const main = async () => {
  console.log("first");
  console.log("second");

  asyncFn(cb)
  // .then(res => {
  //   console.log(res);
  // })
  console.log("last");
  // await prom
  // .then( res => console.log(res)) 
  // .catch(e => console.log(e))
  // .finally(() => console.log('finally')); 
  // console.log("last");

}

main()