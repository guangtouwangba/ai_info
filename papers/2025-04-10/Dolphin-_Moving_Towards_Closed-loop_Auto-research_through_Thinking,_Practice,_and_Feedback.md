# Dolphin: Moving Towards Closed-loop Auto-research through Thinking, Practice, and Feedback

**URL**: http://arxiv.org/abs/2501.03916v3

## 原始摘要

The scientific research paradigm is undergoing a profound transformation
owing to the development of Artificial Intelligence (AI). Recent works
demonstrate that various AI-assisted research methods can largely improve
research efficiency by improving data analysis, accelerating computation, and
fostering novel idea generation. To further move towards the ultimate goal
(i.e., automatic scientific research), in this paper, we introduce Dolphin, a
closed-loop LLM-driven framework to enhance the automation level of scientific
research. Dolphin first generates novel ideas based on feedback from previous
experiments and relevant papers ranked by the topic and task attributes. Then,
the generated ideas can be implemented using a code template refined and
debugged with the designed exception-traceback-guided local code structure.
Finally, Dolphin automatically analyzes the results of each idea and feeds the
results back to the next round of idea generation. Experiments are conducted on
the benchmark datasets of different topics and a subset of MLE-bench. Results
show that Dolphin can continuously improve the performance of the input topic
in a loop. We highlight that Dolphin can automatically propose methods that are
comparable to the state-of-the-art in some tasks such as 3D point
classification.


## AI 摘要

人工智能正在深刻改变科研范式。本文提出的"Dolphin"框架是一个基于大语言模型(LLM)的闭环系统，能自动生成研究想法、实现代码并分析结果。该系统通过分析先前实验反馈和相关论文来产生新想法，使用异常追踪指导的代码结构实现方案，并自动评估结果以改进下一轮研究。在多个基准数据集上的实验表明，Dolphin能持续提升研究性能，在某些任务(如3D点分类)中提出的方法甚至媲美最先进水平。该框架显著提高了科研自动化程度，向"自动科研"目标迈出重要一步。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T12:02:21Z
- **目录日期**: 2025-04-10
