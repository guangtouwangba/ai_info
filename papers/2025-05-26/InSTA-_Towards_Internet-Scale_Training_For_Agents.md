# InSTA: Towards Internet-Scale Training For Agents

**URL**: http://arxiv.org/abs/2502.06776v2

## 原始摘要

The predominant approach for training web navigation agents is to gather
human demonstrations for a set of popular websites and hand-written tasks, but
it is becoming clear that human data is an inefficient resource. We develop a
pipeline to facilitate internet-scale training for agents without laborious
human annotations. In the first stage, an LLM annotates 150k sites with agentic
tasks. In the next stage, LLM agents complete tasks and produce trajectories.
In the final stage, an LLM filters trajectories by judging their success.
Language models are powerful data curation tools, identifying harmful content
with an accuracy of 97%, judging successful trajectories with an accuracy of
82.6%, and producing effective data. We train agents based on Qwen 3 1.7B that
are competitive with frontier LLMs as web agents, while being smaller and
faster. Our top agent reaches a success rate of 56.9%, outperforming the data
collection policy Qwen 3 235B, a 235 times larger Llama 4 Maverick, and
reaching 94.7% of the performance of Gemini 2.5 Flash. We are releasing code,
models and data at: https://data-for-agents.github.io.


## AI 摘要

这篇论文提出了一种无需人工标注的大规模网络导航智能体训练方法。研究者开发了一个三阶段流程：首先用大语言模型(LLM)为15万个网站生成任务，然后让LLM智能体执行任务并生成轨迹，最后用LLM筛选成功轨迹。该方法展现了语言模型强大的数据筛选能力，有害内容识别准确率达97%，轨迹成功判断准确率82.6%。基于Qwen 3 1.7B训练的智能体性能优异，成功率56.9%，超越了235倍大的Llama 4 Maverick，达到Gemini 2.5 Flash性能的94.7%。研究团队将开源代码、模型和数据。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T00:02:37Z
- **目录日期**: 2025-05-26
