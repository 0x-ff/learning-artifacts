'use strict';

let demo = {
    log: function(id, message) {
        console.log('demo ' + id + ' > ' + message);
    }
};

(new Promise((resolve, reject) => {
    demo.log(1, 'Hello, Promise!');
}))
.then(_ => demo.log(1, 'Never'), _ => demo.log(1, 'Displayed'));

(new Promise((resolve, reject) => {
    throw new Error("Banana!");
}))
.catch((e) => demo.log(2, 'Catched: ' + e.message));

(new Promise((resolve, reject) => {
    setTimeout(() => {
      resolve("result");
    }, 400);  
}))
.then(
    result => {
        demo.log(3, "Fulfilled: " + result);
    },
    error => {
        demo.log(3, "Rejected: " + error);
    }
);

(new Promise((resolve, reject) => {
    setTimeout(() => {
      reject(new Error('Opps, I did it again!'));
    }, 400);  
}))
.then(
    result => {
        demo.log(4, "Fulfilled: " + result);
    },
    error => {
        demo.log(4, "Rejected: " + error);
    }
);

(new Promise((resolve, reject) => {
    setTimeout(() => resolve("result"), 300);
    setTimeout(() => reject(new Error("ignored")), 500);  
}))
.then(
    result => demo.log(5, "Fulfilled: " + result),
    error => demo.log(5, "Rejected: " + error)
);

(new Promise((resolve, reject) => {
    resolve('promise 1');
}))
.then(
    function(result) {
        demo.log(6, "Fulfilled: " + result);
        return (new Promise((resolve, reject) => {
            resolve('promise 2');
        }));
    },
    error => demo.log(6, "Rejected: " + error)
)
.then(
    result => demo.log(6, "Fulfilled: " + result),
    error => demo.log(6, "Rejected: " + error)
);
