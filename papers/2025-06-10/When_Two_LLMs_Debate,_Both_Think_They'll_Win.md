# When Two LLMs Debate, Both Think They'll Win

**URL**: http://arxiv.org/abs/2505.19184v3

## 原始摘要

Can LLMs accurately adjust their confidence when facing opposition? Building
on previous studies measuring calibration on static fact-based
question-answering tasks, we evaluate Large Language Models (LLMs) in a
dynamic, adversarial debate setting, uniquely combining two realistic factors:
(a) a multi-turn format requiring models to update beliefs as new information
emerges, and (b) a zero-sum structure to control for task-related uncertainty,
since mutual high-confidence claims imply systematic overconfidence. We
organized 60 three-round policy debates among ten state-of-the-art LLMs, with
models privately rating their confidence (0-100) in winning after each round.
We observed five concerning patterns: (1) Systematic overconfidence: models
began debates with average initial confidence of 72.9% vs. a rational 50%
baseline. (2) Confidence escalation: rather than reducing confidence as debates
progressed, debaters increased their win probabilities, averaging 83% by the
final round. (3) Mutual overestimation: in 61.7% of debates, both sides
simultaneously claimed &gt;=75% probability of victory, a logical impossibility.
(4) Persistent self-debate bias: models debating identical copies increased
confidence from 64.1% to 75.2%; even when explicitly informed their chance of
winning was exactly 50%, confidence still rose (from 50.0% to 57.1%). (5)
Misaligned private reasoning: models' private scratchpad thoughts sometimes
differed from their public confidence ratings, raising concerns about
faithfulness of chain-of-thought reasoning. These results suggest LLMs lack the
ability to accurately self-assess or update their beliefs in dynamic,
multi-turn tasks; a major concern as LLMs are now increasingly deployed without
careful review in assistant and agentic roles.
  Code for our experiments is available at
https://github.com/pradyuprasad/llms_overconfidence


## AI 摘要

该研究评估了大语言模型(LLMs)在动态对抗辩论中的置信度校准问题。研究发现五大问题：(1)系统性过度自信(初始平均72.9% vs 理性基准50%)；(2)辩论中置信度不降反升(最终达83%)；(3)61.7%的辩论出现双方同时声称≥75%胜率的逻辑矛盾；(4)自我辩论时即使告知50%胜率，置信度仍从50%升至57.1%；(5)模型内部推理与公开置信度不一致。结果表明LLMs在动态多轮任务中缺乏准确自我评估和信念更新的能力，这对LLMs在助理和代理角色中的应用构成重大隐患。实验代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T23:02:01Z
- **目录日期**: 2025-06-10
