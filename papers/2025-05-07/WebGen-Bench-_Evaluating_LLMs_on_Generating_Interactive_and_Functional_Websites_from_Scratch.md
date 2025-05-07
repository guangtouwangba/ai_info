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

研究人员提出了WebGen-Bench基准测试，用于评估基于LLM的智能体从零生成多文件网站代码库的能力。该基准包含多样化网站生成指令，涵盖3大类13小类网络应用，并配有647个GPT-4o生成的测试用例。测试采用网页导航代理自动执行，结果显示当前最优组合(Bolt.diy+DeepSeek-R1)准确率仅27.8%。研究还构建了包含6,667条指令的训练集WebGen-Instruct，训练后的Qwen2.5-Coder-32B-Instruct模型准确率达38.2%，超越现有最优专有模型。该基准凸显了网站自动生成任务的挑战性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T19:01:16Z
- **目录日期**: 2025-05-07
