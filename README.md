# TLan

TLan is a console application for time management.

Time management is a topic that can be applied in many contexts, from trivial things such as having a list of tasks for the day, to more profound topics such as how to enjoy life and how to make the best out of time (see list of quotes in the end).

TLan takes a holistic approach to the problem. One starts by defining **purpose**, which speaks to values and goals for a given timeframe. Then **planning** helps with defining the steps to achieve a purpose. **Schedule** is about organizing how things will be executed on daily basis. And, finally, **tracking** is about getting data to make the adjustments necessary to make better use of time.

All these definitions (and more!) are made through a repl. In the repl one can execute commands, like `edit goals`, `goals`, and `schedule`. Commands like `edit goals` will open a vi editor where you can start typing goals (see below). Commands like `goals` and `schedule` will display the contents of the system in a customized UI.

## Model for Time Management

### Purpose
We consider two aspects of purpose in this application: values and goals. Values represent things that are valuable for you, that you want to spend time with, but that doesn't have a specific goal involved. Examples of values are spending time with family and friends, doing meditation, etc. Goals are also valuable but, well, they have a clear goal, i.e., a desired end state. Examples of goals are: to become an AI specialist or moving to a country or becoming a CEO.

In order to define goals, create a file named `goals.gr` and add it to the `data` folder. The layout of this file follows a Domain-Specific Language. We will first show an example and then the definition of this language. 

```
// samples/goals.gr
Goals
- Second bachelors in Mathematics
- Stanford AI certificate
- Write a book
- Have an Intelligent Assistant
``` 

Each of these lines is defining goals. If you want to add high-level themes for the goals, you can add them before the goals in a phrase without the dash (-).
```
// samples/goals.gr
Explore and Build AI 
- BS in Mathematics
- Stanford AI Certificate
- Write a book
- Have an Intelligent Assistant
Write Code and Influence Tech
- Work in many industries
- Be a first-class engineer
- Speak in conferences 
``` 

Once you have defined goals, you can run the application:

```
go run main.go samples

```

What you will see next is the prompt of the repl. As we mentioned before, TLan is a terminal app, and operates through the repl. You can type commands related to all the four domains (i.e, purpose, planning, schedule, and tracking). These commands allow you to do things like editing files and seeing dashboards.

In the context of purpose, you can try the command ```goals```.
```
>> goals
+---------------------------------+-------------------------------+
| EXPLORE AND BUILD AI            | WRITE CODE AND INFLUENCE TECH |
+---------------------------------+-------------------------------+
| Second bachelors in Mathematics | Work in many industries       |
| Stanford AI certificate         | Be a first-class engineer     |
| Write a book                    | Speak in conferences          |
| Have an Intelligent Assistant   |                               |
+---------------------------------+-------------------------------+
```

### Planning
Planning is all about organizing a way to reach your goals. A plan consists of a set to steps taking you from where you are today to the goal. These steps can be either **projects** or **tasks**. Projects are a big set of activities, and they usually take many days or weeks to completes. Tasks are more usually atomic, i.e, things that don't make sense to split anymore.

You can define projects in a file named `projects.gr`:

```
// samples/projects.gr
- Analysis II [01/01-10/05]
- Modern Algebra I [01/01-10/05]
- Number Theory [10/05-31/08]
- Modern Algebra II [10/05-31/08]
- Advanced Mathematics [01/07-31/12]
- CS221 AI [01/04-31/05]
- TLan [01/01-31/03]
```

Note that the dates inside the brackets inform the start and end date for a project. 

You can also group similar projects together under categories, as well as link each project to a goal (`>> goal name`).

```
// samples/projects.gr
Mathematics
- Analysis II [01/01-10/05] >> BS in Mathematics
- Modern Algebra I [01/01-10/05] >> BS in Mathematics
- Number Theory [10/05-31/08] >> BS in Mathematics
- Modern Algebra II [10/05-31/08] >> BS in Mathematics
- Advanced Mathematics [01/09-31/12]

AI
- CS221 AI [01/04-31/05] >> Stanford AI Certificate
- TLan [01/01-31/03] >> Have an Intelligent Assistant
```

You can also add tasks to your projects. Tasks are activities that usually cannot be broken in smaller units.

```
Mathematics
- Analysis II [01/01-10/05] >> BS in Mathematics
  * Read chapter from Burke
  * Read chapter from Lay
  * Do weekly homework
- Modern Algebra I [01/01-10/05] >> BS in Mathematics
- Number Theory [10/05-31/08] >> BS in Mathematics
- Modern Algebra II [10/05-31/08] >> BS in Mathematics
- Advanced Mathematics [01/09-31/12]

AI
- CS221 AI [01/04-31/05] >> Stanford AI Certificate
  * Do weekly homework
- TLan [01/01-31/03] >> Have an Intelligent Assistant

```

Once you have defined projects, you can use the command `plan` to see how your plan for the next few months looks like.

```
+-------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+---------+----------+
|                         | MARCH          | APRIL          | MAY            | JUNE           | JULY           | AUGUST         | SEPTEMBER      | OCTOBER        | NOVEMBER       | DECEMBER       | JANUARY | FEBRUARY |
+-------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+---------+----------+
| Mathematics             |                    Analysis II                   |                  Number Theory                   |                        Advanced Mathematics                       |                    |
|                         +--------------------------------------------------+--------------------------------------------------+-------------------------------------------------------------------+--------------------+
|                         |                 Modern Algebra I                 |                Modern Algebra II                 |                                                                                        |
+-------------------------+----------------+---------------------------------+--------------------------------------------------+----------------------------------------------------------------------------------------+
| AI                      | TLan           |             CS221 AI            |                                                                                                                                           |
+-------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+---------+----------+
```

Now you can call `goals --deep` and this will display not only the goals as we did before, but also the projects associated with those goals.

```
+-------------------------------+-------------------------------+
| EXPLORE AND BUILD AI          | WRITE CODE AND INFLUENCE TECH |
+-------------------------------+-------------------------------+
| BS IN MATHEMATICS             | WORK IN MANY INDUSTRIES       |
| -Analysis II                  |                               |
| -Modern Algebra I             | BE A FIRST CLASS ENGINEER     |
| -Number Theory                |                               |
| -Modern Algebra II            | SPEAK IN CONFERENCES          |
|                               |                               |
| STANFORD AI CERTIFICATE       |                               |
| -CS221 AI                     |                               |
|                               |                               |
| WRITE A BOOK                  |                               |
|                               |                               |
| HAVE AN INTELLIGENT ASSISTANT |                               |
| -TLan                         |                               |
|                               |                               |
+-------------------------------+-------------------------------+
```

### Schedule

Now that you have goals and projects, you can start managing the time that you will be dedicated to these projects in a more practical way. That is why you need a schedule. In a schedule, you separate the days and the time slots that you will be reserving to each project.

The file `schedule.gr` defines your schedule.

```
// samples/schedule.gr
Creative Work [Daily, 05:00-09:00]
- Mathematics [MonFriSatSun]
- AI [TueWedThuSatSun]

Work [Daily, 09:00-18:00]
- Work

Family [Daily, 18:00-22:00]
- Family
```

Each non-dashed line creates a spot on the agenda. You can define the days and times that spot will be valid for. The lines that start with a dash add categories of projects to that time spot. You can have many categories of projects for different days for that spot.

You can see a map of your week by running the command `week`:

```
+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+
|                    | MONDAY             | TUESDAY            | WEDNESDAY          | THURSDAY           | FRIDAY             | SATURDAY           | SUNDAY             |
+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+
| 04:00              |                                                                         x                                                                        |
+--------------------+--------------------+--------------------------------------------------------------+--------------------+-----------------------------------------+
| 05:00              | Mathematics        |                              AI                              | Mathematics        |              Mathematics,AI             |
+--------------------+                    |                                                              |                    |                                         |
| 06:00              |                    |                                                              |                    |                                         |
+--------------------+                    |                                                              |                    |                                         |
| 07:00              |                    |                                                              |                    |                                         |
+--------------------+                    |                                                              |                    |                                         |
| 08:00              |                    |                                                              |                    |                                         |
+--------------------+--------------------+--------------------------------------------------------------+--------------------+-----------------------------------------+
| 09:00              |                                                                       Work                                                                       |
+--------------------+                                                                                                                                                  |
| 10:00              |                                                                                                                                                  |
+--------------------+                                                                                                                                                  |
| 11:00              |                                                                                                                                                  |
+--------------------+                                                                                                                                                  |
| 12:00              |                                                                                                                                                  |
+--------------------+                                                                                                                                                  |
| 13:00              |                                                                                                                                                  |
+--------------------+                                                                                                                                                  |
| 14:00              |                                                                                                                                                  |
+--------------------+                                                                                                                                                  |
| 15:00              |                                                                                                                                                  |
+--------------------+                                                                                                                                                  |
| 16:00              |                                                                                                                                                  |
+--------------------+                                                                                                                                                  |
| 17:00              |                                                                                                                                                  |
+--------------------+--------------------------------------------------------------------------------------------------------------------------------------------------+
| 18:00              |                                                                         x                                                                        |
+--------------------+                                                                                                                                                  |
| 19:00              |                                                                                                                                                  |
+--------------------+--------------------------------------------------------------------------------------------------------------------------------------------------+
| 20:00              |                                                                      Family                                                                      |
+--------------------+                                                                                                                                                  |
| 21:00              |                                                                                                                                                  |
+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+--------------------+
```

Now that you have your schedule, you can also ask TLan what you should be doing now. You do that by running the command `now`:

```
>> now 07:00
NOW is time to do Creative Work

AI
 -- Do weekly homework [CS221 AI]

```

## Quotes
- *"Take care of the minutes and the hours will take care of themselves."* - Lord Chesterfield

- *"I must govern the clock, not be governed by it."* - Golda Meir

- *"Most of us spend too much time on what is urgent, and not enough time on what is important."* - Steven Covey

- *“When you buy something, you’re not paying money for it. You’re paying with the hours of your life you had to spend earning that money. The difference is that life is one thing money can’t buy. Life only gets shorter, and it is pitiful to waste one’s life and freedom that way.”* - José Mujica

- *"Don't be fooled by the calendar. There are only as many days in the year as you make use of. One man gets only a week's value out of a year while another man gets a full year's value out of a week."* - Charles Richards

- *"The common man is not concerned about the passage of time, the man of talent is driven by it."* - Shoppenhauer

- *“I wish it need not have happened in my time,” said Frodo.
“So do I,” said Gandalf, “and so do all who live to see such times. But that is not for them to decide. All we have to decide is what to do with the time that is given us.”* - J.R.R. Tolkien, The Fellowship of the Ring

- *"It's not enough to be busy, so are the ants. The question is, what are we busy about?"* - Henry David Thoreau

- *"Let our advance worrying become advance thinking and planning."* - Winston Churchill

- *"Better to be three hours too soon, than a minute too late."* - William Shakespeare

- *"Yesterday is gone. Tomorrow has not yet come. We have only today. Let us begin."* - Mother Teresa

- *"The shorter way to do many things is to only do one thing at a time."* - Mozart

- *"Give me six hours to chop down a tree and I will spend the first four sharpening the axe."* - Abraham Lincoln

- *"Determine never to be idle. No person will have occasion to complain of the want of time who never loses any. It is wonderful how much can be done if we are always doing."* - Thomas Jefferson

- *"Once you have mastered time, you will understand how true it is that most people overestimate what they can accomplish in a year - and underestimate what they can achieve in a decade!"* -- Tony Robbins

- *“The bad news is time flies. The good news is you’re the pilot”* - Michael Altshuler

- *“You may delay, but time will not.”* - Benjamin Franklin

- *“A man who dares to waste one hour of life has not discovered the value of life.”* - Charles Darwin


### Tracking
[TBD]

## Input Files
[TBD]

## Commands
[TBD]
