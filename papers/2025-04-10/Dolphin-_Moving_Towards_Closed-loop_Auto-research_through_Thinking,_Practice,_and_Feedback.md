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

人工智能（AI）正在推动科研范式的变革。本文提出的Dolphin框架是一个闭环的LLM驱动系统，旨在提升科研自动化水平。它通过分析先前实验反馈和相关论文生成新想法，利用调试优化的代码模板实现方案，并自动分析结果以指导下一轮研究。实验表明，Dolphin能在循环中持续提升目标课题性能，在某些任务（如3D点分类）中自动提出的方法甚至媲美当前最优成果。该框架通过整合AI辅助的数据分析、计算加速和创新生成，向自动化科研的终极目标迈进了一步。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T05:02:15Z
- **目录日期**: 2025-04-10
