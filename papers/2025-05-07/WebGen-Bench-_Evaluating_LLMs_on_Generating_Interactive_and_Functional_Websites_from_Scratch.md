# WebGen-Bench: Evaluating LLMs on Generating Interactive and Functional Websites from Scratch

**URL**: http://arxiv.org/abs/2505.03733v1

## 原始摘要

LLM-based agents have demonstrated great potential in generating and managing
code within complex codebases. In this paper, we introduce WebGen-Bench, a
novel benchmark designed to measure an LLM-based agent's ability to create
multi-file website codebases from scratch. It contains diverse instructions for
website generation, created through the combined efforts of human annotators
and GPT-4o. These instructions span three major categories and thirteen minor
categories, encompassing nearly all important types of web applications. To
assess the quality of the generated websites, we use GPT-4o to generate test
cases targeting each functionality described in the instructions, and then
manually filter, adjust, and organize them to ensure accuracy, resulting in 647
test cases. Each test case specifies an operation to be performed on the
website and the expected result after the operation. To automate testing and
improve reproducibility, we employ a powerful web-navigation agent to execute
tests on the generated websites and determine whether the observed responses
align with the expected results. We evaluate three high-performance code-agent
frameworks, Bolt.diy, OpenHands, and Aider, using multiple proprietary and
open-source LLMs as engines. The best-performing combination, Bolt.diy powered
by DeepSeek-R1, achieves only 27.8\% accuracy on the test cases, highlighting
the challenging nature of our benchmark. Additionally, we construct
WebGen-Instruct, a training set consisting of 6,667 website-generation
instructions. Training Qwen2.5-Coder-32B-Instruct on Bolt.diy trajectories
generated from a subset of this training set achieves an accuracy of 38.2\%,
surpassing the performance of the best proprietary model.


## AI 摘要

本文介绍了WebGen-Bench，一个评估基于LLM的智能体从零创建多文件网站代码库能力的新基准。该基准包含由人类标注员和GPT-4o共同生成的多样化网站生成指令，涵盖三大类和十三小类，几乎覆盖所有重要Web应用类型。研究人员使用GPT-4o生成647个测试用例，并通过人工筛选确保准确性。测试采用强大的网页导航代理自动执行。评估显示，性能最佳的组合（Bolt.diy + DeepSeek-R1）仅达到27.8%的准确率。此外，基于6,667条指令构建的训练集WebGen-Instruct训练出的模型Qwen2.5-Coder-32B-Instruct达到了38.2%的准确率，超越了最佳专有模型。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T04:01:23Z
- **目录日期**: 2025-05-07
