# Boosting Rust Unit Test Coverage through Hybrid Program Analysis and Large Language Models

**URL**: http://arxiv.org/abs/2506.09002v1

## 原始摘要

Unit testing is essential for ensuring software reliability and correctness.
Classic Search-Based Software Testing (SBST) methods and concolic
execution-based approaches for generating unit tests often fail to achieve high
coverage due to difficulties in handling complex program units, such as
branching conditions and external dependencies. Recent work has increasingly
utilized large language models (LLMs) to generate test cases, improving the
quality of test generation by providing better context and correcting errors in
the model's output. However, these methods rely on fixed prompts, resulting in
relatively low compilation success rates and coverage. This paper presents
PALM, an approach that leverages large language models (LLMs) to enhance the
generation of high-coverage unit tests. PALM performs program analysis to
identify branching conditions within functions, which are then combined into
path constraints. These constraints and relevant contextual information are
used to construct prompts that guide the LLMs in generating unit tests. We
implement the approach and evaluate it in 10 open-source Rust crates.
Experimental results show that within just two or three hours, PALM can
significantly improves test coverage compared to classic methods, with
increases in overall project coverage exceeding 50% in some instances and its
generated tests achieving an average coverage of 75.77%, comparable to human
effort (71.30%), highlighting the potential of LLMs in automated test
generation. We submitted 91 PALM-generated unit tests targeting new code. Of
these submissions, 80 were accepted, 5 were rejected, and 6 remain pending
review. The results demonstrate the effectiveness of integrating program
analysis with AI and open new avenues for future research in automated software
testing.


## AI 摘要

传统基于搜索的软件测试(SBST)和符号执行方法在生成单元测试时难以处理复杂程序单元，导致覆盖率较低。本研究提出PALM方法，通过程序分析识别函数分支条件并构建路径约束，结合上下文信息生成提示词指导大语言模型(LLM)生成测试用例。在10个Rust开源项目中的实验表明，PALM仅需2-3小时即可显著提升测试覆盖率，某些项目覆盖率提升超50%，平均覆盖率达75.77%，接近人工水平(71.30%)。提交的91个测试中有80个被接受，验证了该方法在自动化测试中的有效性，为AI与程序分析结合开辟了新途径。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T05:02:45Z
- **目录日期**: 2025-06-11
