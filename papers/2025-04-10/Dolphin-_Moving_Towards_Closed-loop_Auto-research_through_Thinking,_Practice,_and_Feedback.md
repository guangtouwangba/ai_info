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

人工智能正推动科研范式的变革。本文提出Dolphin框架，通过闭环式大语言模型驱动实现科研自动化：1）基于实验反馈和相关文献生成新思路；2）通过异常回溯引导的代码模板实现方案；3）自动分析结果并反馈优化。在多个基准测试中，Dolphin能持续提升研究主题表现，在3D点分类等任务中提出的方法甚至达到最先进水平。该框架通过"生成-实现-反馈"的闭环机制，显著提升了从数据分析到创新想法生成的科研全流程自动化程度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T18:02:10Z
- **目录日期**: 2025-04-10
