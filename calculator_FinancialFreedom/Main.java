import java.text.NumberFormat;

public class Main {
    public static void main(String[] args) {
        
        // USER DEFINED
        long initial_Investment = 1_000;
        float salary = 5000f;
        float monthlySalaryShare = 0.20f; // How much the user invests, as a % of the salary
        float yield = 0.03f;
        short investment_years = 10;
        String operation = "Net Worth"; // possible values: "Salary Share", "Yield", "Time", "Net Worth"
        

        // final byte MONTHS_IN_YEARS = 12;
        // double needed_netWorth = salary*MONTHS_IN_YEARS/yield; // The amount of money the user wishes to have at the end



        


        switch (operation){
            case "Net Worth":
                CompoundProjector netWorthObj = new CompoundProjector(initial_Investment, salary, monthlySalaryShare, yield, investment_years);


                double result_end_Investment = netWorthObj.projectNetWorth();

                double monthlyDividends = netWorthObj.projectDividends();

                System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(result_end_Investment));
                System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(monthlyDividends));
                break;

            case "Yield":

                CompoundProjector yieldOBJ = new CompoundProjector(initial_Investment, salary, monthlySalaryShare, yield, investment_years);
                yieldOBJ.setInitial_Investment(initial_Investment);
                yieldOBJ.setSalary(salary);
                yieldOBJ.setMonthlySalaryShare(monthlySalaryShare);
                yieldOBJ.setYield(yield);
                yieldOBJ.setInvestment_years(investment_years);
                
                //System.out.println("Estimated Yield: " + guessYield);
                break;
            case "Time":
                break;
            case "Salary Share":
                break;
        }

        




        
    }
}
