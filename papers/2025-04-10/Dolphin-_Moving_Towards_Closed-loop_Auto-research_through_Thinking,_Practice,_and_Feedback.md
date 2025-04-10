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

人工智能正在推动科研范式的变革。本文提出的"Dolphin"是一个闭环式大语言模型驱动框架，旨在提升科研自动化水平：1）基于先前实验反馈和相关论文生成新想法；2）通过异常回溯引导的代码结构实现想法；3）自动分析结果并反馈优化。实验表明，Dolphin能在循环中持续提升输入主题的性能，在3D点分类等任务中提出的方法甚至可比肩当前最优方案。该框架通过自动化的"生成-实现-反馈"闭环，显著提高了数据处理、计算加速和新想法产生的效率，向全自动科研目标迈出重要一步。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T06:03:15Z
- **目录日期**: 2025-04-10
