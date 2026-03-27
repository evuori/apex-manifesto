# AI Engagement Guide
### Using AI Tools Effectively Within APEX

---

## The Core Principle

Under APEX, the rule about AI is simple: **The expert stays in charge.** AI is an accelerant. It produces output. Experts produce judgment.

You are the owner of every output produced in service of your work, regardless of how it was generated. If the AI generated it wrong, that is your problem — not the AI's. There is no "the AI made me do it."

This guide translates that principle into actual practice: when to use AI, how to structure that use, and what to watch for.

---

## The Decision Framework: Should I Use AI for This?

Ask yourself three questions in order. Answer them honestly.

### Question 1: Do I Understand This Problem?

Can you, without AI, evaluate whether the output is correct?

- **Yes** → Go to Question 2
- **No, I do not understand it** → Do not use AI on this task right now. Take time to understand the problem. Understanding first, AI acceleration second.
- **Yes, but it is tedious** → Go to Question 2
- **Yes, but it would take a long time** → Go to Question 2

If you cannot independently evaluate the output, you cannot responsibly ship it. Using AI to substitute for understanding is not efficiency — it is risk without awareness.

### Question 2: Would I Do This Task If AI Did Not Exist?

Is this actually your work, or are you outsourcing because it is easier?

- **Yes, I would do it** → Go to Question 3
- **No, I would skip it or do something else** → Do not use AI. If it is not worth your time, it probably is not worth doing.

This filters out the temptation to use AI to generate busywork. The question is: is this a real task that serves your delivery? If yes, you might accelerate it. If no, skip it.

### Question 3: Where Is the Judgment?

Where does domain expertise matter in this task?

- **All over the place** (defining the approach, evaluating options, deciding on trade-offs, assessing quality) → AI handles maybe 30-40% of the work. Use it for the routine parts, keep the judgment for yourself.
- **At the start** (setting up the problem, deciding the approach) → Do the judgment work first, then give AI clear constraints to work within.
- **At the end** (evaluating and refining output) → Use AI to generate options quickly. You do the evaluation.
- **At the start and end** (setup and evaluation) → Use AI for the middle repetitive work.
- **Nowhere — this is just producing output** → Be careful. Tasks with no judgment are rare. Usually there is judgment you are not seeing.

Now you have a rough map of where AI fits. Use it accordingly.

---

## Practical Patterns

### Pattern 1: Code Generation

**Judgment at start:** Define the function contract, outline the logic, set constraints.
**Judgment at end:** Review the generated code. Does it handle edge cases? Is it efficient? Does it match your style?

**How to do this:**
1. Write the function signature and a clear docstring explaining what it does and what it returns
2. Outline the approach in a comment (3-4 sentences)
3. Ask AI to implement it
4. Read the code. If you do not understand it or it looks wrong, ask the AI to change it or write it yourself
5. Test it

**What goes wrong:**
- Asking AI to write a function without any specification
- Accepting the first output without reading it
- Testing it only for the happy path
- Using the code without understanding what it does

**Red flag:** You are using AI for a function you could not write yourself. That is not acceleration — that is delegation. Stop and write it yourself first, or reassign the task.

---

### Pattern 2: Documentation and Communication

**Judgment at start:** Decide what matters (structure, key points, tone)
**Judgment at end:** Does the output say what you meant? Is it accurate?

**How to do this:**
1. Outline the document or message: what are the key points? What order? What tone?
2. Give AI the outline and ask it to draft the full text
3. Read it. Edit it. Make sure it says what you meant.
4. If you find yourself heavily editing it, rewrite it yourself — the drafting is not saving you time

**What goes wrong:**
- Using AI output without reading or editing (your credibility, not the AI's, is on the line)
- Asking for a document without an outline (you get generic output that sounds like an AI wrote it)
- Accepting AI's organizational choices (they are not always right for your audience)
- Shipping something with factual errors because you did not verify

**Red flag:** You are using an AI-drafted document without substantial editing. That means the AI understood your intent better than you can articulate it. Rewrite from scratch to understand what you are trying to say.

---

### Pattern 3: Analysis and Research

**Judgment at start:** Define what you are trying to understand. What questions need answering?
**Judgment at end:** Is the analysis actually correct? Does it answer the question?

**How to do this:**
1. Tell AI: "I am trying to understand X. Here is the data/context. What patterns do you see?"
2. Read the analysis. Is it consistent with the data you know? Does it make sense?
3. If the analysis finds something surprising, verify it independently before acting on it
4. Use the output to inform your thinking, not replace it

**What goes wrong:**
- Accepting surprising conclusions without independent verification
- Assuming the AI has access to all relevant context (it usually does not)
- Not checking whether the AI is hallucinating or making assumptions
- Using analysis without understanding the reasoning

**Red flag:** You are making a significant decision based on AI-generated analysis that you have not independently verified. Stop and verify it first.

---

### Pattern 4: Refinement and Iteration

**Judgment throughout:** You are directing the iterations.

**How to do this:**
1. You have a draft or a piece of work that is close but not quite right
2. Ask AI to refine it in a specific way: "Make this more concise," "Add more examples," "Simplify the technical language"
3. Review. If it is better, keep it. If not, revert or edit further.
4. Repeat until you are satisfied

**What goes wrong:**
- Asking for generic "improvements" without specifying what would actually improve it
- Iterating 10+ times (if you have iterated more than 3 times, you probably need to rewrite it yourself)
- Losing your own voice in the output as you iterate
- Treating iteration as a process that ends when the AI stops suggesting changes

**Red flag:** You have iterated more than 3 times and the output is still not right. The AI is not understanding what you want. Write it yourself or clarify your intent more precisely.

---

### Pattern 5: Exploration and Options

**Judgment at start:** Define what you are exploring.
**Judgment at end:** Choose among options.

**How to do this:**
1. Ask AI: "Generate 3-5 different approaches to this problem. For each, outline the pros and cons."
2. Read the options. Do they make sense? Are they realistic?
3. Evaluate them against your constraints (time, complexity, maintainability, etc.)
4. Choose the one that best fits your situation

**What goes wrong:**
- Treating the first option as the default
- Accepting the AI's pros/cons uncritically (evaluate them yourself)
- Not considering whether all the options are actually viable in your context
- Using this pattern for decisions that require deep domain knowledge (you should be generating the options)

**Red flag:** The AI's options are more creative or better thought-out than what you would have come up with. This usually means you have not spent enough time thinking about the problem. Think more. Then use AI to explore further.

---

## Anti-Patterns: What Goes Wrong

### Anti-Pattern 1: Using AI Because Something Is Hard

**Red flag:** "This is too complicated, let me ask AI."

If something is hard because you do not understand it, using AI does not fix it. You still do not understand it. You just have output you cannot evaluate.

If something is hard because it is genuinely tedious (write 20 similar functions, format data, etc.), then AI is appropriate.

---

### Anti-Pattern 2: Accepting Output Without Review

**Red flag:** "The AI generated it, so it must be right."

AI is confident. Confident output is not the same as correct output. The AI can generate plausible-looking code that does not work, correct-sounding analysis that is wrong, or professional-looking documents with factual errors.

You are responsible. Review what you ship.

---

### Anti-Pattern 3: Using AI to Avoid Necessary Learning

**Red flag:** "I don't need to learn this language / framework / domain because AI can do it for me."

This fails as soon as the AI output breaks or needs to be adjusted. You cannot debug or improve code you don't understand. You cannot evaluate analysis of a domain you have not studied. You cannot make good decisions about things you do not understand.

AI is for acceleration, not replacement of knowledge.

---

### Anti-Pattern 4: Treating AI as a Person

**Red flag:** "The AI said X, so it must be true" or "I asked the AI three times and it kept saying the same thing, so I trusted it."

The AI is not a person. It does not have special knowledge or judgment. It can be wrong consistently and confidently. Verification is your responsibility.

---

### Anti-Pattern 5: Over-Relying on AI for Core Work

**Red flag:** Most of your core work is now AI-generated. You are primarily refining AI output rather than creating.

This is brittle. You lose the ability to evaluate your own work. You lose the context of your decisions. You become dependent on the AI working well.

Keep a healthy balance: you should be creating maybe 40-60% of your output directly, accelerating maybe 30-40% with AI, and minimizing the rest.

---

### Anti-Pattern 6: Unclear Ownership

**Red flag:** "I used AI for that" is presented as an explanation rather than context.

Using AI is context. It does not change who is accountable for the output. Under APEX, you own it completely.

If you are invoking AI as an excuse ("the AI made a mistake, not me"), you are in anti-pattern territory.

---

## Standards for Different Outputs

### Code

Standard: You can read and understand it. You can debug it. You can modify it. Test it. Even if you used AI to write it, you must be able to own it technically.

**Yes use AI:** Boilerplate, implementations of well-defined algorithms, translations between similar patterns
**No:** Core logic you do not understand, systems-level decisions, anything security-critical without expert review

### Documentation

Standard: It says what you mean. It is accurate. It represents your project fairly.

**Yes use AI:** Drafting, structuring, expanding outlines, refining tone
**No:** Core messages, important claims, anything you have not personally verified

### Analysis

Standard: Conclusions are verified. You understand the reasoning. The output answers your actual question.

**Yes use AI:** Finding patterns, organizing data, generating hypotheses to test
**No:** Final conclusions, decisions affecting your roadmap, anything that informs the direction of your project without verification

### Communication (email, updates, etc.)

Standard: It sounds like you. It says what you want to say. It represents your judgment.

**Yes use AI:** Drafting routine communications, expanding a bullet list into prose
**No:** Anything client-facing without review, communications about sensitive topics, anything you don't fully agree with

---

## The Test

If you cannot articulate why you chose a tool or approach, and explain your reasoning to someone else, you are not ready to ship it. This is true with or without AI.

AI is a tool for people who know what they are doing. It is dangerous for people who are unclear. Be honest about which category you are in right now.

---

*Author: Erno Vuori <erno.vuori@gmail.com>*
*Repository: https://github.com/evuori/apex-manifesto*
