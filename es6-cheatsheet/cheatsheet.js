
//
// let 
let apples = 5; // (*)
if (true) {
  let apples = 10;
  alert(apples); // 10 (внутри блока)
}
alert(apples); // 5 (снаружи блока значение не изменилось)

// 
// const
const apple = 5;
apple = 10; // ошибка

const user = {
    name: "Вася"
};
user.name = "Петя"; // допустимо
user = 5; // нельзя, будет ошибка

//
// destructuring assignment
let [firstName, lastName] = ["Илья", "Кантор"];
let [, , title] = "Юлий Цезарь Император Рима".split(" ");
let [firstName="Гость", lastName="Анонимный"] = [];
let [firstName, lastName, ...rest] = 
    "Юлий Цезарь Император Рима".split(" "); // spread

let options = {
    title: "Меню",
    width: 100,
    height: 200
};
let {title, width, height} = options;
let {width: w, height: h, title} = options;
let {width=100, height=200, title} = options;

let options = {
    size: {
      width: 100,
      height: 200
    },
    items: ["Пончик", "Пирожное"]
}
let { title="Меню", size: {width, height}, items: [item1, item2] } = options;

//
// functions
function showMenu(title = "Без заголовка", width = 100, height = 200) {}
function showName(firstName, lastName, ...rest) {}
function showMenu({title, width, height}) {}
function showMenu({title="Заголовок", width:w=100, height:h=200}) {}
let inc = x => x+1; // function(x) { return x+1; }
let sum = (a,b) => a + b;
let getTime = () => new Date().getHours() + ':' + new Date().getMinutes();

let group = {
    title: "Наш курс",
    students: ["Вася", "Петя", "Даша"],
    showList: function() {
      this.students.forEach(
        student => alert(this.title + ': ' + student)
      )
    }
}
group.showList();

// 
// strings
let str = `обратные кавычки`;
alert(`моя
  многострочная
  строка`);
let apples = 2;
let oranges = 3;
alert(`${apples} + ${oranges} = ${apples + oranges}`); // 2 + 3 = 5
let str = func`моя строка`;

function f(strings, ...values) {
    alert(JSON.stringify(strings));     // ["Sum of "," + "," =\n ","!"]
    alert(JSON.stringify(strings.raw)); // ["Sum of "," + "," =\\n ","!"]
    alert(JSON.stringify(values));      // [3,5,8]
}  
let apples = 3;
let oranges = 5;
//          |  s[0] | v[0] |s[1]| v[1]  |s[2]  |      v[2]      |s[3]
let str = f`Sum of ${apples} + ${oranges} =\n ${apples + oranges}!`;

// 
// objects
let name = "Вася";
let isAdmin = true;
let user = {name,isAdmin}; // {"name": "Вася", "isAdmin": true}

let propName = "firstName";
let user = {[propName]: "Вася"}; // user.firstName

let name = "Вася";
let user = {
  name,
  // вместо "sayHi: function() {...}" пишем "sayHi() {...}"
  sayHi() {
    alert(this.name);
  }
};

let methodName = "getFirstName";
let user = {
  // в квадратных скобках может быть любое выражение,
  // которое должно вернуть название метода
  [methodName]() {  // вместо [methodName]: function() {
    return "Вася";
  }
};

let animal = {
    walk() {
        alert("I'm walking");
    }
};
let rabbit = {
    __proto__: animal,
    walk() {
        alert(super.walk); // walk() { … }
        super.walk(); // I'm walking
    }
};
rabbit.walk();

// 
// class
/*
class Название [extends Родитель]  {
    constructor
    методы
}
*/
class User {
    constructor(name) {
        this.name = name;
    }
    sayHi() {
        alert(this.name);
    }
}
let user = new User("Вася");
user.sayHi(); // Вася

let User = class {
   sayHi() { alert('Привет!'); }
};  
new User().sayHi();

class User {
    constructor(firstName, lastName) {
      this.firstName = firstName;
      this.lastName = lastName;
    }
  
    static createGuest() {
      return new User("Гость", "Сайта");
    }
};  
let user = User.createGuest();

class Animal {
    constructor(name) {
      this.name = name;
    }
  
    walk() {
      alert("I walk: " + this.name);
    }
}  
class Rabbit extends Animal {
    walk() {
      super.walk();
      alert("...and jump!");
    }
}
new Rabbit("Вася").walk();