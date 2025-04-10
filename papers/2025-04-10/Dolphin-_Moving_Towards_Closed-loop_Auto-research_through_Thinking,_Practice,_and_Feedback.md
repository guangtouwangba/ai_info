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

人工智能（AI）正在推动科研范式的变革。本文提出Dolphin框架，利用大语言模型（LLM）实现闭环自动化科研：1）基于实验反馈和相关文献生成新想法；2）通过异常回溯引导的代码模板实现想法；3）自动分析结果并反馈优化。实验表明，Dolphin能在循环中持续提升课题性能，在3D点分类等任务中提出的方法甚至媲美最先进技术。该框架通过自动化的数据分析和计算加速，显著提高了科研效率，为实现全自动科研迈出重要一步。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T14:02:25Z
- **目录日期**: 2025-04-10
