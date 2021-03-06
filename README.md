# TLan

TLan is a console application for time management.

Time management is a topic that can be applied in many contexts, from trivial things such as having a list of tasks for the day, to more profound topics such as how to enjoy life and how to make the best out of time (see list of quotes in the end).

TLan takes a holistic approach to the problem. One starts by defining **purpose**, which speaks to values and goals for a given timeframe. Then **planning** helps with defining the steps to get there. **Schedule** is about organizing how things will be executed on daily basis. And, finally, **tracking** is about getting data to make the adjustments necessary to make a better use of time.

## Model for Time Management

### Purpose
We consider two aspects of purpose in this application: values and goals. Values represent things that are valuable for you, that you want to spend time with, but that don't a specific goal involved. Examples of values are spending time with family and friend, doing meditation, etc. Goals are also valuable but, well, they have a clear goal, an end state that you want to reach. Examples goals are: to become an AI specialist or moving to a country or becoming a CEO.

In order to define goals, create a file `goals.gr` and add it to the `data` folder. The layout of this file follows a Domain-Specific Language. We will first show an example and then the definition of this language. 

```
// samples/goals.gr
Become a nonfiction novelist
Influence the future of Artificial Intelligence
Live in at least five different cultures
Improve body health
``` 

This is a file that is defining high-level goals. These goals should be very aspirational, and not necessarily yet need to be detailed. If you want to add more details, you can add sub-goals in the following way.
```
// samples/goals.gr
Become a nonfiction novelist
- Write first book on fake news
Influence the future of AI
Live in at least five different cultures
Improve body health
- Establish a gym rhytm
- Lose weight
``` 

Once you have defined goals, you can run the application:

```
go run main.go samples
```

What you will see next is the prompt of the repl. As we mentioned before, TLan is a terminal app, and operates through the repl. You can type commands related to all the four domains (i.e, purpose, planning, schedule and tracking). These commands allow you to do things like editing files and seeing dashboards.

In the context of purpose, you can try the command ```goals```.
```
>> goals
+----------------------------+------------------------------------------+-------------------------------+-----------------------+
| INFLUENCE THE FUTURE OF AI | LIVE IN AT LEAST FIVE DIFFERENT CULTURES | BECOME A NONFICTION NOVELIST  | IMPROVE BODY HEALTH   |
+----------------------------+------------------------------------------+-------------------------------+-----------------------+
| Influence the future of AI | Live in at least five different cultures | Write first book on fake news | Establish a gym rhytm |
|                            |                                          |                               | Lose weight           |
+----------------------------+------------------------------------------+-------------------------------+-----------------------+
```

### Planning

### Schedule

### Tracking

## Input Files

## Commands

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